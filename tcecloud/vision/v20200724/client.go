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

package v20200724

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-07-24"

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

func NewDeleteProdCollRequest() (request *DeleteProdCollRequest) {
	request = &DeleteProdCollRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "DeleteProdColl")
	return
}

func NewDeleteProdCollResponse() (response *DeleteProdCollResponse) {
	response = &DeleteProdCollResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除数据集
func (c *Client) DeleteProdColl(request *DeleteProdCollRequest) (response *DeleteProdCollResponse, err error) {
	if request == nil {
		request = NewDeleteProdCollRequest()
	}
	response = NewDeleteProdCollResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlakeTimerLogsRequest() (request *DescribeFlakeTimerLogsRequest) {
	request = &DescribeFlakeTimerLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "DescribeFlakeTimerLogs")
	return
}

func NewDescribeFlakeTimerLogsResponse() (response *DescribeFlakeTimerLogsResponse) {
	response = &DescribeFlakeTimerLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询触发器执行历史
func (c *Client) DescribeFlakeTimerLogs(request *DescribeFlakeTimerLogsRequest) (response *DescribeFlakeTimerLogsResponse, err error) {
	if request == nil {
		request = NewDescribeFlakeTimerLogsRequest()
	}
	response = NewDescribeFlakeTimerLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFlakeViewRequest() (request *DeleteFlakeViewRequest) {
	request = &DeleteFlakeViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "DeleteFlakeView")
	return
}

func NewDeleteFlakeViewResponse() (response *DeleteFlakeViewResponse) {
	response = &DeleteFlakeViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除视图
func (c *Client) DeleteFlakeView(request *DeleteFlakeViewRequest) (response *DeleteFlakeViewResponse, err error) {
	if request == nil {
		request = NewDeleteFlakeViewRequest()
	}
	response = NewDeleteFlakeViewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFlakeViewRequest() (request *CreateFlakeViewRequest) {
	request = &CreateFlakeViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "CreateFlakeView")
	return
}

func NewCreateFlakeViewResponse() (response *CreateFlakeViewResponse) {
	response = &CreateFlakeViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建视图
func (c *Client) CreateFlakeView(request *CreateFlakeViewRequest) (response *CreateFlakeViewResponse, err error) {
	if request == nil {
		request = NewCreateFlakeViewRequest()
	}
	response = NewCreateFlakeViewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFlakeTimerRequest() (request *DeleteFlakeTimerRequest) {
	request = &DeleteFlakeTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "DeleteFlakeTimer")
	return
}

func NewDeleteFlakeTimerResponse() (response *DeleteFlakeTimerResponse) {
	response = &DeleteFlakeTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除触发器
func (c *Client) DeleteFlakeTimer(request *DeleteFlakeTimerRequest) (response *DeleteFlakeTimerResponse, err error) {
	if request == nil {
		request = NewDeleteFlakeTimerRequest()
	}
	response = NewDeleteFlakeTimerResponse()
	err = c.Send(request, response)
	return
}

func NewVisionRequest() (request *VisionRequest) {
	request = &VisionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "Vision")
	return
}

func NewVisionResponse() (response *VisionResponse) {
	response = &VisionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端资源概览接口
func (c *Client) Vision(request *VisionRequest) (response *VisionResponse, err error) {
	if request == nil {
		request = NewVisionRequest()
	}
	response = NewVisionResponse()
	err = c.Send(request, response)
	return
}

func NewOspVisionRequest() (request *OspVisionRequest) {
	request = &OspVisionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "OspVision")
	return
}

func NewOspVisionResponse() (response *OspVisionResponse) {
	response = &OspVisionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运维端资源概览接口
func (c *Client) OspVision(request *OspVisionRequest) (response *OspVisionResponse, err error) {
	if request == nil {
		request = NewOspVisionRequest()
	}
	response = NewOspVisionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFlakeTimerRequest() (request *CreateFlakeTimerRequest) {
	request = &CreateFlakeTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "CreateFlakeTimer")
	return
}

func NewCreateFlakeTimerResponse() (response *CreateFlakeTimerResponse) {
	response = &CreateFlakeTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建触发器
func (c *Client) CreateFlakeTimer(request *CreateFlakeTimerRequest) (response *CreateFlakeTimerResponse, err error) {
	if request == nil {
		request = NewCreateFlakeTimerRequest()
	}
	response = NewCreateFlakeTimerResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProdCollRequest() (request *CreateProdCollRequest) {
	request = &CreateProdCollRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "CreateProdColl")
	return
}

func NewCreateProdCollResponse() (response *CreateProdCollResponse) {
	response = &CreateProdCollResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建数据集基本信息
func (c *Client) CreateProdColl(request *CreateProdCollRequest) (response *CreateProdCollResponse, err error) {
	if request == nil {
		request = NewCreateProdCollRequest()
	}
	response = NewCreateProdCollResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProdCollsRequest() (request *DescribeProdCollsRequest) {
	request = &DescribeProdCollsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "DescribeProdColls")
	return
}

func NewDescribeProdCollsResponse() (response *DescribeProdCollsResponse) {
	response = &DescribeProdCollsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照条件查询数据集，返回列表
func (c *Client) DescribeProdColls(request *DescribeProdCollsRequest) (response *DescribeProdCollsResponse, err error) {
	if request == nil {
		request = NewDescribeProdCollsRequest()
	}
	response = NewDescribeProdCollsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlakeViewsRequest() (request *DescribeFlakeViewsRequest) {
	request = &DescribeFlakeViewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "DescribeFlakeViews")
	return
}

func NewDescribeFlakeViewsResponse() (response *DescribeFlakeViewsResponse) {
	response = &DescribeFlakeViewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询视图列表
func (c *Client) DescribeFlakeViews(request *DescribeFlakeViewsRequest) (response *DescribeFlakeViewsResponse, err error) {
	if request == nil {
		request = NewDescribeFlakeViewsRequest()
	}
	response = NewDescribeFlakeViewsResponse()
	err = c.Send(request, response)
	return
}

func NewCallFlakeViewRequest() (request *CallFlakeViewRequest) {
	request = &CallFlakeViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "CallFlakeView")
	return
}

func NewCallFlakeViewResponse() (response *CallFlakeViewResponse) {
	response = &CallFlakeViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用视图
func (c *Client) CallFlakeView(request *CallFlakeViewRequest) (response *CallFlakeViewResponse, err error) {
	if request == nil {
		request = NewCallFlakeViewRequest()
	}
	response = NewCallFlakeViewResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateFlakeTimerRequest() (request *UpdateFlakeTimerRequest) {
	request = &UpdateFlakeTimerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "UpdateFlakeTimer")
	return
}

func NewUpdateFlakeTimerResponse() (response *UpdateFlakeTimerResponse) {
	response = &UpdateFlakeTimerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改触发器
func (c *Client) UpdateFlakeTimer(request *UpdateFlakeTimerRequest) (response *UpdateFlakeTimerResponse, err error) {
	if request == nil {
		request = NewUpdateFlakeTimerRequest()
	}
	response = NewUpdateFlakeTimerResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateProdCollRequest() (request *UpdateProdCollRequest) {
	request = &UpdateProdCollRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "UpdateProdColl")
	return
}

func NewUpdateProdCollResponse() (response *UpdateProdCollResponse) {
	response = &UpdateProdCollResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改，更新数据集信息
func (c *Client) UpdateProdColl(request *UpdateProdCollRequest) (response *UpdateProdCollResponse, err error) {
	if request == nil {
		request = NewUpdateProdCollRequest()
	}
	response = NewUpdateProdCollResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlakeTimersRequest() (request *DescribeFlakeTimersRequest) {
	request = &DescribeFlakeTimersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "DescribeFlakeTimers")
	return
}

func NewDescribeFlakeTimersResponse() (response *DescribeFlakeTimersResponse) {
	response = &DescribeFlakeTimersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询触发器
func (c *Client) DescribeFlakeTimers(request *DescribeFlakeTimersRequest) (response *DescribeFlakeTimersResponse, err error) {
	if request == nil {
		request = NewDescribeFlakeTimersRequest()
	}
	response = NewDescribeFlakeTimersResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateFlakeViewRequest() (request *UpdateFlakeViewRequest) {
	request = &UpdateFlakeViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("vision", APIVersion, "UpdateFlakeView")
	return
}

func NewUpdateFlakeViewResponse() (response *UpdateFlakeViewResponse) {
	response = &UpdateFlakeViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改视图
func (c *Client) UpdateFlakeView(request *UpdateFlakeViewRequest) (response *UpdateFlakeViewResponse, err error) {
	if request == nil {
		request = NewUpdateFlakeViewRequest()
	}
	response = NewUpdateFlakeViewResponse()
	err = c.Send(request, response)
	return
}
