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

package v20200119

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-01-19"

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

func NewGetRegionZoneIdcInfoRequest() (request *GetRegionZoneIdcInfoRequest) {
	request = &GetRegionZoneIdcInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("gconf", APIVersion, "GetRegionZoneIdcInfo")
	return
}

func NewGetRegionZoneIdcInfoResponse() (response *GetRegionZoneIdcInfoResponse) {
	response = &GetRegionZoneIdcInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询地域可用区idc信息
func (c *Client) GetRegionZoneIdcInfo(request *GetRegionZoneIdcInfoRequest) (response *GetRegionZoneIdcInfoResponse, err error) {
	if request == nil {
		request = NewGetRegionZoneIdcInfoRequest()
	}
	response = NewGetRegionZoneIdcInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetZonesInfoRequest() (request *GetZonesInfoRequest) {
	request = &GetZonesInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("gconf", APIVersion, "GetZonesInfo")
	return
}

func NewGetZonesInfoResponse() (response *GetZonesInfoResponse) {
	response = &GetZonesInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取可用区信息
func (c *Client) GetZonesInfo(request *GetZonesInfoRequest) (response *GetZonesInfoResponse, err error) {
	if request == nil {
		request = NewGetZonesInfoRequest()
	}
	response = NewGetZonesInfoResponse()
	err = c.Send(request, response)
	return
}
