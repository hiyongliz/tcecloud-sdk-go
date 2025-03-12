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

package v20221013

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-10-13"

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

func NewDeletePvgwGwRequest() (request *DeletePvgwGwRequest) {
	request = &DeletePvgwGwRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "DeletePvgwGw")
	return
}

func NewDeletePvgwGwResponse() (response *DeletePvgwGwResponse) {
	response = &DeletePvgwGwResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除PVGW2.0网关集群
func (c *Client) DeletePvgwGw(request *DeletePvgwGwRequest) (response *DeletePvgwGwResponse, err error) {
	if request == nil {
		request = NewDeletePvgwGwRequest()
	}
	response = NewDeletePvgwGwResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEipRequest() (request *DescribeEipRequest) {
	request = &DescribeEipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "DescribeEip")
	return
}

func NewDescribeEipResponse() (response *DescribeEipResponse) {
	response = &DescribeEipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeEip）用于查询EIP信息
func (c *Client) DescribeEip(request *DescribeEipRequest) (response *DescribeEipResponse, err error) {
	if request == nil {
		request = NewDescribeEipRequest()
	}
	response = NewDescribeEipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatDnatRequest() (request *DescribeNatDnatRequest) {
	request = &DescribeNatDnatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "DescribeNatDnat")
	return
}

func NewDescribeNatDnatResponse() (response *DescribeNatDnatResponse) {
	response = &DescribeNatDnatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取nat的dnat规则
func (c *Client) DescribeNatDnat(request *DescribeNatDnatRequest) (response *DescribeNatDnatResponse, err error) {
	if request == nil {
		request = NewDescribeNatDnatRequest()
	}
	response = NewDescribeNatDnatResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateNatGatewayRequest() (request *MigrateNatGatewayRequest) {
	request = &MigrateNatGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "MigrateNatGateway")
	return
}

func NewMigrateNatGatewayResponse() (response *MigrateNatGatewayResponse) {
	response = &MigrateNatGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NAT网关实例迁移
func (c *Client) MigrateNatGateway(request *MigrateNatGatewayRequest) (response *MigrateNatGatewayResponse, err error) {
	if request == nil {
		request = NewMigrateNatGatewayRequest()
	}
	response = NewMigrateNatGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePvgwGwRequest() (request *CreatePvgwGwRequest) {
	request = &CreatePvgwGwRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "CreatePvgwGw")
	return
}

func NewCreatePvgwGwResponse() (response *CreatePvgwGwResponse) {
	response = &CreatePvgwGwResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建PVGW2.0网关集群
func (c *Client) CreatePvgwGw(request *CreatePvgwGwRequest) (response *CreatePvgwGwResponse, err error) {
	if request == nil {
		request = NewCreatePvgwGwRequest()
	}
	response = NewCreatePvgwGwResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatSnatRequest() (request *DescribeNatSnatRequest) {
	request = &DescribeNatSnatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "DescribeNatSnat")
	return
}

func NewDescribeNatSnatResponse() (response *DescribeNatSnatResponse) {
	response = &DescribeNatSnatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取nat的snat规则
func (c *Client) DescribeNatSnat(request *DescribeNatSnatRequest) (response *DescribeNatSnatResponse, err error) {
	if request == nil {
		request = NewDescribeNatSnatRequest()
	}
	response = NewDescribeNatSnatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePvgwGwRequest() (request *DescribePvgwGwRequest) {
	request = &DescribePvgwGwRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "DescribePvgwGw")
	return
}

func NewDescribePvgwGwResponse() (response *DescribePvgwGwResponse) {
	response = &DescribePvgwGwResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取PVGW网关信息
func (c *Client) DescribePvgwGw(request *DescribePvgwGwRequest) (response *DescribePvgwGwResponse, err error) {
	if request == nil {
		request = NewDescribePvgwGwRequest()
	}
	response = NewDescribePvgwGwResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointIpv6Request() (request *DescribeVpcEndPointIpv6Request) {
	request = &DescribeVpcEndPointIpv6Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "DescribeVpcEndPointIpv6")
	return
}

func NewDescribeVpcEndPointIpv6Response() (response *DescribeVpcEndPointIpv6Response) {
	response = &DescribeVpcEndPointIpv6Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Ipv6终端连接信息
func (c *Client) DescribeVpcEndPointIpv6(request *DescribeVpcEndPointIpv6Request) (response *DescribeVpcEndPointIpv6Response, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointIpv6Request()
	}
	response = NewDescribeVpcEndPointIpv6Response()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointServiceIpv6Request() (request *DescribeVpcEndPointServiceIpv6Request) {
	request = &DescribeVpcEndPointServiceIpv6Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "DescribeVpcEndPointServiceIpv6")
	return
}

func NewDescribeVpcEndPointServiceIpv6Response() (response *DescribeVpcEndPointServiceIpv6Response) {
	response = &DescribeVpcEndPointServiceIpv6Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Ipv6终端连接服务信息
func (c *Client) DescribeVpcEndPointServiceIpv6(request *DescribeVpcEndPointServiceIpv6Request) (response *DescribeVpcEndPointServiceIpv6Response, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointServiceIpv6Request()
	}
	response = NewDescribeVpcEndPointServiceIpv6Response()
	err = c.Send(request, response)
	return
}

func NewUpdatePvgwGwRequest() (request *UpdatePvgwGwRequest) {
	request = &UpdatePvgwGwRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dfsvpc", APIVersion, "UpdatePvgwGw")
	return
}

func NewUpdatePvgwGwResponse() (response *UpdatePvgwGwResponse) {
	response = &UpdatePvgwGwResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新PVGW2.0网关集群
func (c *Client) UpdatePvgwGw(request *UpdatePvgwGwRequest) (response *UpdatePvgwGwResponse, err error) {
	if request == nil {
		request = NewUpdatePvgwGwRequest()
	}
	response = NewUpdatePvgwGwResponse()
	err = c.Send(request, response)
	return
}
