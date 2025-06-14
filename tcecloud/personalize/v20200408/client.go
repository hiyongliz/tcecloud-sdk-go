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

package v20200408

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-04-08"

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

func NewGetRequest() (request *GetRequest) {
	request = &GetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("personalize", APIVersion, "Get")
	return
}

func NewGetResponse() (response *GetResponse) {
	response = &GetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户个性化配置
func (c *Client) Get(request *GetRequest) (response *GetResponse, err error) {
	if request == nil {
		request = NewGetRequest()
	}
	response = NewGetResponse()
	err = c.Send(request, response)
	return
}

func NewPutRequest() (request *PutRequest) {
	request = &PutRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("personalize", APIVersion, "Put")
	return
}

func NewPutResponse() (response *PutResponse) {
	response = &PutResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑用户个性化设置
func (c *Client) Put(request *PutRequest) (response *PutResponse, err error) {
	if request == nil {
		request = NewPutRequest()
	}
	response = NewPutResponse()
	err = c.Send(request, response)
	return
}
