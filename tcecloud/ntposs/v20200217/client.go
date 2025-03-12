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

func NewGetMonitorDataRequest() (request *GetMonitorDataRequest) {
	request = &GetMonitorDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ntposs", APIVersion, "GetMonitorData")
	return
}

func NewGetMonitorDataResponse() (response *GetMonitorDataResponse) {
	response = &GetMonitorDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ntp监控数据
func (c *Client) GetMonitorData(request *GetMonitorDataRequest) (response *GetMonitorDataResponse, err error) {
	if request == nil {
		request = NewGetMonitorDataRequest()
	}
	response = NewGetMonitorDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetOpInfoListRequest() (request *GetOpInfoListRequest) {
	request = &GetOpInfoListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ntposs", APIVersion, "GetOpInfoList")
	return
}

func NewGetOpInfoListResponse() (response *GetOpInfoListResponse) {
	response = &GetOpInfoListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作记录筛选列表
func (c *Client) GetOpInfoList(request *GetOpInfoListRequest) (response *GetOpInfoListResponse, err error) {
	if request == nil {
		request = NewGetOpInfoListRequest()
	}
	response = NewGetOpInfoListResponse()
	err = c.Send(request, response)
	return
}

func NewOperateNtpRequest() (request *OperateNtpRequest) {
	request = &OperateNtpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ntposs", APIVersion, "OperateNtp")
	return
}

func NewOperateNtpResponse() (response *OperateNtpResponse) {
	response = &OperateNtpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交ntp操作
func (c *Client) OperateNtp(request *OperateNtpRequest) (response *OperateNtpResponse, err error) {
	if request == nil {
		request = NewOperateNtpRequest()
	}
	response = NewOperateNtpResponse()
	err = c.Send(request, response)
	return
}

func NewUploadClientTimeStampRequest() (request *UploadClientTimeStampRequest) {
	request = &UploadClientTimeStampRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ntposs", APIVersion, "UploadClientTimeStamp")
	return
}

func NewUploadClientTimeStampResponse() (response *UploadClientTimeStampResponse) {
	response = &UploadClientTimeStampResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传客户端时间(此接口已废弃)
func (c *Client) UploadClientTimeStamp(request *UploadClientTimeStampRequest) (response *UploadClientTimeStampResponse, err error) {
	if request == nil {
		request = NewUploadClientTimeStampRequest()
	}
	response = NewUploadClientTimeStampResponse()
	err = c.Send(request, response)
	return
}

func NewGetNtpConfRequest() (request *GetNtpConfRequest) {
	request = &GetNtpConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ntposs", APIVersion, "GetNtpConf")
	return
}

func NewGetNtpConfResponse() (response *GetNtpConfResponse) {
	response = &GetNtpConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取ntp配置信息
func (c *Client) GetNtpConf(request *GetNtpConfRequest) (response *GetNtpConfResponse, err error) {
	if request == nil {
		request = NewGetNtpConfRequest()
	}
	response = NewGetNtpConfResponse()
	err = c.Send(request, response)
	return
}

func NewGetOpRecordsRequest() (request *GetOpRecordsRequest) {
	request = &GetOpRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ntposs", APIVersion, "GetOpRecords")
	return
}

func NewGetOpRecordsResponse() (response *GetOpRecordsResponse) {
	response = &GetOpRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作记录
func (c *Client) GetOpRecords(request *GetOpRecordsRequest) (response *GetOpRecordsResponse, err error) {
	if request == nil {
		request = NewGetOpRecordsRequest()
	}
	response = NewGetOpRecordsResponse()
	err = c.Send(request, response)
	return
}
