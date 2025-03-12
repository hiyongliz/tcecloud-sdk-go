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

package v20200217

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-02-17"

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

func NewQueryCertificateApplyListRequest() (request *QueryCertificateApplyListRequest) {
	request = &QueryCertificateApplyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("certify", APIVersion, "QueryCertificateApplyList")
	return
}

func NewQueryCertificateApplyListResponse() (response *QueryCertificateApplyListResponse) {
	response = &QueryCertificateApplyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取审核信息列表
func (c *Client) QueryCertificateApplyList(request *QueryCertificateApplyListRequest) (response *QueryCertificateApplyListResponse, err error) {
	if request == nil {
		request = NewQueryCertificateApplyListRequest()
	}
	response = NewQueryCertificateApplyListResponse()
	err = c.Send(request, response)
	return
}

func NewAllowCertificateApplyRequest() (request *AllowCertificateApplyRequest) {
	request = &AllowCertificateApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("certify", APIVersion, "AllowCertificateApply")
	return
}

func NewAllowCertificateApplyResponse() (response *AllowCertificateApplyResponse) {
	response = &AllowCertificateApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审核通过
func (c *Client) AllowCertificateApply(request *AllowCertificateApplyRequest) (response *AllowCertificateApplyResponse, err error) {
	if request == nil {
		request = NewAllowCertificateApplyRequest()
	}
	response = NewAllowCertificateApplyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCertificateStatusRequest() (request *UpdateCertificateStatusRequest) {
	request = &UpdateCertificateStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("certify", APIVersion, "UpdateCertificateStatus")
	return
}

func NewUpdateCertificateStatusResponse() (response *UpdateCertificateStatusResponse) {
	response = &UpdateCertificateStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更资质审核状态
func (c *Client) UpdateCertificateStatus(request *UpdateCertificateStatusRequest) (response *UpdateCertificateStatusResponse, err error) {
	if request == nil {
		request = NewUpdateCertificateStatusRequest()
	}
	response = NewUpdateCertificateStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCertificateStatusRequest() (request *QueryCertificateStatusRequest) {
	request = &QueryCertificateStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("certify", APIVersion, "QueryCertificateStatus")
	return
}

func NewQueryCertificateStatusResponse() (response *QueryCertificateStatusResponse) {
	response = &QueryCertificateStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前资质审核状态
func (c *Client) QueryCertificateStatus(request *QueryCertificateStatusRequest) (response *QueryCertificateStatusResponse, err error) {
	if request == nil {
		request = NewQueryCertificateStatusRequest()
	}
	response = NewQueryCertificateStatusResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProvinceAndCityInfoRequest() (request *QueryProvinceAndCityInfoRequest) {
	request = &QueryProvinceAndCityInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("certify", APIVersion, "QueryProvinceAndCityInfo")
	return
}

func NewQueryProvinceAndCityInfoResponse() (response *QueryProvinceAndCityInfoResponse) {
	response = &QueryProvinceAndCityInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取省份和城市信息
func (c *Client) QueryProvinceAndCityInfo(request *QueryProvinceAndCityInfoRequest) (response *QueryProvinceAndCityInfoResponse, err error) {
	if request == nil {
		request = NewQueryProvinceAndCityInfoRequest()
	}
	response = NewQueryProvinceAndCityInfoResponse()
	err = c.Send(request, response)
	return
}

func NewRefuseCertificateApplyRequest() (request *RefuseCertificateApplyRequest) {
	request = &RefuseCertificateApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("certify", APIVersion, "RefuseCertificateApply")
	return
}

func NewRefuseCertificateApplyResponse() (response *RefuseCertificateApplyResponse) {
	response = &RefuseCertificateApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审核拒绝
func (c *Client) RefuseCertificateApply(request *RefuseCertificateApplyRequest) (response *RefuseCertificateApplyResponse, err error) {
	if request == nil {
		request = NewRefuseCertificateApplyRequest()
	}
	response = NewRefuseCertificateApplyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCertificateDetailRequest() (request *QueryCertificateDetailRequest) {
	request = &QueryCertificateDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("certify", APIVersion, "QueryCertificateDetail")
	return
}

func NewQueryCertificateDetailResponse() (response *QueryCertificateDetailResponse) {
	response = &QueryCertificateDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取审核详情
func (c *Client) QueryCertificateDetail(request *QueryCertificateDetailRequest) (response *QueryCertificateDetailResponse, err error) {
	if request == nil {
		request = NewQueryCertificateDetailRequest()
	}
	response = NewQueryCertificateDetailResponse()
	err = c.Send(request, response)
	return
}
