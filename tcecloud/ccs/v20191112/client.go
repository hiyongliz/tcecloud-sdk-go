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

package v20191112

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2019-11-12"

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

func NewGetTopKUserDataRequest() (request *GetTopKUserDataRequest) {
	request = &GetTopKUserDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetTopKUserData")
	return
}

func NewGetTopKUserDataResponse() (response *GetTopKUserDataResponse) {
	response = &GetTopKUserDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetTopKUserData
func (c *Client) GetTopKUserData(request *GetTopKUserDataRequest) (response *GetTopKUserDataResponse, err error) {
	if request == nil {
		request = NewGetTopKUserDataRequest()
	}
	response = NewGetTopKUserDataResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterQuotaRequest() (request *ModifyClusterQuotaRequest) {
	request = &ModifyClusterQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "ModifyClusterQuota")
	return
}

func NewModifyClusterQuotaResponse() (response *ModifyClusterQuotaResponse) {
	response = &ModifyClusterQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群和单集群节点配额
func (c *Client) ModifyClusterQuota(request *ModifyClusterQuotaRequest) (response *ModifyClusterQuotaResponse, err error) {
	if request == nil {
		request = NewModifyClusterQuotaRequest()
	}
	response = NewModifyClusterQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewModifyQuotaRequest() (request *ModifyQuotaRequest) {
	request = &ModifyQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "ModifyQuota")
	return
}

func NewModifyQuotaResponse() (response *ModifyQuotaResponse) {
	response = &ModifyQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改ccr配额
func (c *Client) ModifyQuota(request *ModifyQuotaRequest) (response *ModifyQuotaResponse, err error) {
	if request == nil {
		request = NewModifyQuotaRequest()
	}
	response = NewModifyQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewGetK8sVersionDistributionRequest() (request *GetK8sVersionDistributionRequest) {
	request = &GetK8sVersionDistributionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetK8sVersionDistribution")
	return
}

func NewGetK8sVersionDistributionResponse() (response *GetK8sVersionDistributionResponse) {
	response = &GetK8sVersionDistributionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetK8sVersionDistribution
func (c *Client) GetK8sVersionDistribution(request *GetK8sVersionDistributionRequest) (response *GetK8sVersionDistributionResponse, err error) {
	if request == nil {
		request = NewGetK8sVersionDistributionRequest()
	}
	response = NewGetK8sVersionDistributionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "DescribeCluster")
	return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群列表。
func (c *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRequest()
	}
	response = NewDescribeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopKReturnUserDataRequest() (request *GetTopKReturnUserDataRequest) {
	request = &GetTopKReturnUserDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetTopKReturnUserData")
	return
}

func NewGetTopKReturnUserDataResponse() (response *GetTopKReturnUserDataResponse) {
	response = &GetTopKReturnUserDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetTopKReturnUserData
func (c *Client) GetTopKReturnUserData(request *GetTopKReturnUserDataRequest) (response *GetTopKReturnUserDataResponse, err error) {
	if request == nil {
		request = NewGetTopKReturnUserDataRequest()
	}
	response = NewGetTopKReturnUserDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetStatisticsDataRequest() (request *GetStatisticsDataRequest) {
	request = &GetStatisticsDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetStatisticsData")
	return
}

func NewGetStatisticsDataResponse() (response *GetStatisticsDataResponse) {
	response = &GetStatisticsDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetStatisticsData
func (c *Client) GetStatisticsData(request *GetStatisticsDataRequest) (response *GetStatisticsDataResponse, err error) {
	if request == nil {
		request = NewGetStatisticsDataRequest()
	}
	response = NewGetStatisticsDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopKGrowUserDataRequest() (request *GetTopKGrowUserDataRequest) {
	request = &GetTopKGrowUserDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetTopKGrowUserData")
	return
}

func NewGetTopKGrowUserDataResponse() (response *GetTopKGrowUserDataResponse) {
	response = &GetTopKGrowUserDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetTopKGrowUserData
func (c *Client) GetTopKGrowUserData(request *GetTopKGrowUserDataRequest) (response *GetTopKGrowUserDataResponse, err error) {
	if request == nil {
		request = NewGetTopKGrowUserDataRequest()
	}
	response = NewGetTopKGrowUserDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceRequest() (request *DescribeInstanceRequest) {
	request = &DescribeInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "DescribeInstance")
	return
}

func NewDescribeInstanceResponse() (response *DescribeInstanceResponse) {
	response = &DescribeInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群节点列表
func (c *Client) DescribeInstance(request *DescribeInstanceRequest) (response *DescribeInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceRequest()
	}
	response = NewDescribeInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetQuotaRequest() (request *GetQuotaRequest) {
	request = &GetQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetQuota")
	return
}

func NewGetQuotaResponse() (response *GetQuotaResponse) {
	response = &GetQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ccr配额
func (c *Client) GetQuota(request *GetQuotaRequest) (response *GetQuotaResponse, err error) {
	if request == nil {
		request = NewGetQuotaRequest()
	}
	response = NewGetQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewGetDailyStatByAppIdRequest() (request *GetDailyStatByAppIdRequest) {
	request = &GetDailyStatByAppIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetDailyStatByAppId")
	return
}

func NewGetDailyStatByAppIdResponse() (response *GetDailyStatByAppIdResponse) {
	response = &GetDailyStatByAppIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetDailyStatByAppId
func (c *Client) GetDailyStatByAppId(request *GetDailyStatByAppIdRequest) (response *GetDailyStatByAppIdResponse, err error) {
	if request == nil {
		request = NewGetDailyStatByAppIdRequest()
	}
	response = NewGetDailyStatByAppIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterQuotaRequest() (request *GetClusterQuotaRequest) {
	request = &GetClusterQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "GetClusterQuota")
	return
}

func NewGetClusterQuotaResponse() (response *GetClusterQuotaResponse) {
	response = &GetClusterQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群配额
func (c *Client) GetClusterQuota(request *GetClusterQuotaRequest) (response *GetClusterQuotaResponse, err error) {
	if request == nil {
		request = NewGetClusterQuotaRequest()
	}
	response = NewGetClusterQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewTkeApiRequest() (request *TkeApiRequest) {
	request = &TkeApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ccs", APIVersion, "TkeApi")
	return
}

func NewTkeApiResponse() (response *TkeApiResponse) {
	response = &TkeApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TkeApi
func (c *Client) TkeApi(request *TkeApiRequest) (response *TkeApiResponse, err error) {
	if request == nil {
		request = NewTkeApiRequest()
	}
	response = NewTkeApiResponse()
	err = c.Send(request, response)
	return
}
