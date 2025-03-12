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

package v20200105

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-01-05"

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

func NewGetRealTimeClockRequest() (request *GetRealTimeClockRequest) {
	request = &GetRealTimeClockRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("license", APIVersion, "GetRealTimeClock")
	return
}

func NewGetRealTimeClockResponse() (response *GetRealTimeClockResponse) {
	response = &GetRealTimeClockResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务端实时时钟，  响应时间采用RFC3339 格式化
func (c *Client) GetRealTimeClock(request *GetRealTimeClockRequest) (response *GetRealTimeClockResponse, err error) {
	if request == nil {
		request = NewGetRealTimeClockRequest()
	}
	response = NewGetRealTimeClockResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseStatRequest() (request *DescribeLicenseStatRequest) {
	request = &DescribeLicenseStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("license", APIVersion, "DescribeLicenseStat")
	return
}

func NewDescribeLicenseStatResponse() (response *DescribeLicenseStatResponse) {
	response = &DescribeLicenseStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权统计信息
func (c *Client) DescribeLicenseStat(request *DescribeLicenseStatRequest) (response *DescribeLicenseStatResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseStatRequest()
	}
	response = NewDescribeLicenseStatResponse()
	err = c.Send(request, response)
	return
}

func NewLicenseUpdateRequest() (request *LicenseUpdateRequest) {
	request = &LicenseUpdateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("license", APIVersion, "LicenseUpdate")
	return
}

func NewLicenseUpdateResponse() (response *LicenseUpdateResponse) {
	response = &LicenseUpdateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新客户授权
func (c *Client) LicenseUpdate(request *LicenseUpdateRequest) (response *LicenseUpdateResponse, err error) {
	if request == nil {
		request = NewLicenseUpdateRequest()
	}
	response = NewLicenseUpdateResponse()
	err = c.Send(request, response)
	return
}

func NewListProductUsagesRequest() (request *ListProductUsagesRequest) {
	request = &ListProductUsagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("license", APIVersion, "ListProductUsages")
	return
}

func NewListProductUsagesResponse() (response *ListProductUsagesResponse) {
	response = &ListProductUsagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品用量列表
func (c *Client) ListProductUsages(request *ListProductUsagesRequest) (response *ListProductUsagesResponse, err error) {
	if request == nil {
		request = NewListProductUsagesRequest()
	}
	response = NewListProductUsagesResponse()
	err = c.Send(request, response)
	return
}

func NewLicenseDetailRequest() (request *LicenseDetailRequest) {
	request = &LicenseDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("license", APIVersion, "LicenseDetail")
	return
}

func NewLicenseDetailResponse() (response *LicenseDetailResponse) {
	response = &LicenseDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询授权和用量信息
func (c *Client) LicenseDetail(request *LicenseDetailRequest) (response *LicenseDetailResponse, err error) {
	if request == nil {
		request = NewLicenseDetailRequest()
	}
	response = NewLicenseDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetProductEditionRequest() (request *GetProductEditionRequest) {
	request = &GetProductEditionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("license", APIVersion, "GetProductEdition")
	return
}

func NewGetProductEditionResponse() (response *GetProductEditionResponse) {
	response = &GetProductEditionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询软件版本授权: 基础版/高级版/铂金版
func (c *Client) GetProductEdition(request *GetProductEditionRequest) (response *GetProductEditionResponse, err error) {
	if request == nil {
		request = NewGetProductEditionRequest()
	}
	response = NewGetProductEditionResponse()
	err = c.Send(request, response)
	return
}
