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

package v20200512

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-05-12"

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

func NewQueryCreateProductFailureMigrateTaskDetailRequest() (request *QueryCreateProductFailureMigrateTaskDetailRequest) {
	request = &QueryCreateProductFailureMigrateTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "QueryCreateProductFailureMigrateTaskDetail")
	return
}

func NewQueryCreateProductFailureMigrateTaskDetailResponse() (response *QueryCreateProductFailureMigrateTaskDetailResponse) {
	response = &QueryCreateProductFailureMigrateTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询产品故障转移结果
func (c *Client) QueryCreateProductFailureMigrateTaskDetail(request *QueryCreateProductFailureMigrateTaskDetailRequest) (response *QueryCreateProductFailureMigrateTaskDetailResponse, err error) {
	if request == nil {
		request = NewQueryCreateProductFailureMigrateTaskDetailRequest()
	}
	response = NewQueryCreateProductFailureMigrateTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductDataReplicationStateRequest() (request *QueryProductDataReplicationStateRequest) {
	request = &QueryProductDataReplicationStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "QueryProductDataReplicationState")
	return
}

func NewQueryProductDataReplicationStateResponse() (response *QueryProductDataReplicationStateResponse) {
	response = &QueryProductDataReplicationStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询产品实例数据同步状态信息
func (c *Client) QueryProductDataReplicationState(request *QueryProductDataReplicationStateRequest) (response *QueryProductDataReplicationStateResponse, err error) {
	if request == nil {
		request = NewQueryProductDataReplicationStateRequest()
	}
	response = NewQueryProductDataReplicationStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceSpecificationsRequest() (request *DescribeDeviceSpecificationsRequest) {
	request = &DescribeDeviceSpecificationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeDeviceSpecifications")
	return
}

func NewDescribeDeviceSpecificationsResponse() (response *DescribeDeviceSpecificationsResponse) {
	response = &DescribeDeviceSpecificationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询设备规格
func (c *Client) DescribeDeviceSpecifications(request *DescribeDeviceSpecificationsRequest) (response *DescribeDeviceSpecificationsResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceSpecificationsRequest()
	}
	response = NewDescribeDeviceSpecificationsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
	request = &DescribeDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeDevices")
	return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
	response = &DescribeDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询设备信息
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
	if request == nil {
		request = NewDescribeDevicesRequest()
	}
	response = NewDescribeDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductStateInfoRequest() (request *QueryProductStateInfoRequest) {
	request = &QueryProductStateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "QueryProductStateInfo")
	return
}

func NewQueryProductStateInfoResponse() (response *QueryProductStateInfoResponse) {
	response = &QueryProductStateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询产品集群状态和节点的分布情况
func (c *Client) QueryProductStateInfo(request *QueryProductStateInfoRequest) (response *QueryProductStateInfoResponse, err error) {
	if request == nil {
		request = NewQueryProductStateInfoRequest()
	}
	response = NewQueryProductStateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductHealthStateRequest() (request *QueryProductHealthStateRequest) {
	request = &QueryProductHealthStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "QueryProductHealthState")
	return
}

func NewQueryProductHealthStateResponse() (response *QueryProductHealthStateResponse) {
	response = &QueryProductHealthStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询产品健康状态，结果中应该包盖该产品一个或者多个维度的健康状态和总体健康状态。
func (c *Client) QueryProductHealthState(request *QueryProductHealthStateRequest) (response *QueryProductHealthStateResponse, err error) {
	if request == nil {
		request = NewQueryProductHealthStateRequest()
	}
	response = NewQueryProductHealthStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例信息
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductFailureMigrateRequest() (request *CreateProductFailureMigrateRequest) {
	request = &CreateProductFailureMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "CreateProductFailureMigrate")
	return
}

func NewCreateProductFailureMigrateResponse() (response *CreateProductFailureMigrateResponse) {
	response = &CreateProductFailureMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询产品实例数据同步状态信息
func (c *Client) CreateProductFailureMigrate(request *CreateProductFailureMigrateRequest) (response *CreateProductFailureMigrateResponse, err error) {
	if request == nil {
		request = NewCreateProductFailureMigrateRequest()
	}
	response = NewCreateProductFailureMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
	request = &DescribeClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeClusters")
	return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
	response = &DescribeClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群总览
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
	if request == nil {
		request = NewDescribeClustersRequest()
	}
	response = NewDescribeClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCapacityStatisticsRequest() (request *DescribeCapacityStatisticsRequest) {
	request = &DescribeCapacityStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("optdsql", APIVersion, "DescribeCapacityStatistics")
	return
}

func NewDescribeCapacityStatisticsResponse() (response *DescribeCapacityStatisticsResponse) {
	response = &DescribeCapacityStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容量统计信息
func (c *Client) DescribeCapacityStatistics(request *DescribeCapacityStatisticsRequest) (response *DescribeCapacityStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeCapacityStatisticsRequest()
	}
	response = NewDescribeCapacityStatisticsResponse()
	err = c.Send(request, response)
	return
}
