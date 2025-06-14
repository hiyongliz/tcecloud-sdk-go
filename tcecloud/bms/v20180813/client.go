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

package v20180813

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-08-13"

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

func NewDescribeBmsDeployTasksRequest() (request *DescribeBmsDeployTasksRequest) {
	request = &DescribeBmsDeployTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bms", APIVersion, "DescribeBmsDeployTasks")
	return
}

func NewDescribeBmsDeployTasksResponse() (response *DescribeBmsDeployTasksResponse) {
	response = &DescribeBmsDeployTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取BMS生产部署的任务详情
func (c *Client) DescribeBmsDeployTasks(request *DescribeBmsDeployTasksRequest) (response *DescribeBmsDeployTasksResponse, err error) {
	if request == nil {
		request = NewDescribeBmsDeployTasksRequest()
	}
	response = NewDescribeBmsDeployTasksResponse()
	err = c.Send(request, response)
	return
}
