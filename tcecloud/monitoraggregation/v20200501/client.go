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

package v20200501

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-05-01"

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

func NewGetAggreDataRequest() (request *GetAggreDataRequest) {
	request = &GetAggreDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("monitoraggregation", APIVersion, "GetAggreData")
	return
}

func NewGetAggreDataResponse() (response *GetAggreDataResponse) {
	response = &GetAggreDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 聚合argus和cmdb的数据
func (c *Client) GetAggreData(request *GetAggreDataRequest) (response *GetAggreDataResponse, err error) {
	if request == nil {
		request = NewGetAggreDataRequest()
	}
	response = NewGetAggreDataResponse()
	err = c.Send(request, response)
	return
}
