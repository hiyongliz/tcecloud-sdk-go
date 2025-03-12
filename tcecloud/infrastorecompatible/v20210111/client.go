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

package v20210111

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2021-01-11"

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

func NewGetDataBaradMetricRequest() (request *GetDataBaradMetricRequest) {
	request = &GetDataBaradMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastorecompatible", APIVersion, "GetDataBaradMetric")
	return
}

func NewGetDataBaradMetricResponse() (response *GetDataBaradMetricResponse) {
	response = &GetDataBaradMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取barad格式的监控数据
func (c *Client) GetDataBaradMetric(request *GetDataBaradMetricRequest) (response *GetDataBaradMetricResponse, err error) {
	if request == nil {
		request = NewGetDataBaradMetricRequest()
	}
	response = NewGetDataBaradMetricResponse()
	err = c.Send(request, response)
	return
}

func NewGetMonitorDataV2018Request() (request *GetMonitorDataV2018Request) {
	request = &GetMonitorDataV2018Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastorecompatible", APIVersion, "GetMonitorDataV2018")
	return
}

func NewGetMonitorDataV2018Response() (response *GetMonitorDataV2018Response) {
	response = &GetMonitorDataV2018Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取云产品的监控数据。传入产品的命名空间、对象维度描述和监控指标即可获得相应的监控数据。
// 接口调用频率限制为：20次/秒，1200次/分钟。
// 若您需要调用的指标、对象较多，可能存在因限频出现拉取失败的情况，建议尽量将请求按时间维度均摊。
func (c *Client) GetMonitorDataV2018(request *GetMonitorDataV2018Request) (response *GetMonitorDataV2018Response, err error) {
	if request == nil {
		request = NewGetMonitorDataV2018Request()
	}
	response = NewGetMonitorDataV2018Response()
	err = c.Send(request, response)
	return
}

func NewDescribeProductEventListRequest() (request *DescribeProductEventListRequest) {
	request = &DescribeProductEventListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("infrastorecompatible", APIVersion, "DescribeProductEventList")
	return
}

func NewDescribeProductEventListResponse() (response *DescribeProductEventListResponse) {
	response = &DescribeProductEventListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页获取产品事件的列表
func (c *Client) DescribeProductEventList(request *DescribeProductEventListRequest) (response *DescribeProductEventListResponse, err error) {
	if request == nil {
		request = NewDescribeProductEventListRequest()
	}
	response = NewDescribeProductEventListResponse()
	err = c.Send(request, response)
	return
}
