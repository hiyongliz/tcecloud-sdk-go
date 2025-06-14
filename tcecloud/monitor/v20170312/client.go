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

package v20170312

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2017-03-12"

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

func NewDescribeBaseMetricsRequest() (request *DescribeBaseMetricsRequest) {
	request = &DescribeBaseMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeBaseMetrics")
	return
}

func NewDescribeBaseMetricsResponse() (response *DescribeBaseMetricsResponse) {
	response = &DescribeBaseMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取基础指标详情
func (c *Client) DescribeBaseMetrics(request *DescribeBaseMetricsRequest) (response *DescribeBaseMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeBaseMetricsRequest()
	}
	response = NewDescribeBaseMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmMetricsRequest() (request *DescribeAlarmMetricsRequest) {
	request = &DescribeAlarmMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "DescribeAlarmMetrics")
	return
}

func NewDescribeAlarmMetricsResponse() (response *DescribeAlarmMetricsResponse) {
	response = &DescribeAlarmMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用于配置告警的指标列表
func (c *Client) DescribeAlarmMetrics(request *DescribeAlarmMetricsRequest) (response *DescribeAlarmMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmMetricsRequest()
	}
	response = NewDescribeAlarmMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
	request = &GetMonitorDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitor", APIVersion, "GetMonitorData")
	return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
	response = &GetMonitorDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云产品的监控数据。传入产品的命名空间、对象维度描述和监控指标即可获得相应的监控数据。
// 接口调用频率限制为：50次/秒，500次/分钟。
// 若您需要调用的指标、对象较多，可能存在因限频出现拉取失败的情况，建议尽量将请求按时间维度均摊。
func (c *Client) GetMonitorData(request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
	if request == nil {
		request = NewGetMonitorDataRequest()
	}
	response = NewGetMonitorDataResponse()
	err = c.Send(request, response)
	return
}
