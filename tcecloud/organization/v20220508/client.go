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

package v20220508

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2022-05-08"

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

func NewDescribeOrganizationsRequest() (request *DescribeOrganizationsRequest) {
	request = &DescribeOrganizationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizations")
	return
}

func NewDescribeOrganizationsResponse() (response *DescribeOrganizationsResponse) {
	response = &DescribeOrganizationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集团组织列表
func (c *Client) DescribeOrganizations(request *DescribeOrganizationsRequest) (response *DescribeOrganizationsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationsRequest()
	}
	response = NewDescribeOrganizationsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationRequest() (request *DeleteOrganizationRequest) {
	request = &DeleteOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganization")
	return
}

func NewDeleteOrganizationResponse() (response *DeleteOrganizationResponse) {
	response = &DeleteOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集团组织
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationRequest()
	}
	response = NewDeleteOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationMemberRequest() (request *CreateOrganizationMemberRequest) {
	request = &CreateOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMember")
	return
}

func NewCreateOrganizationMemberResponse() (response *CreateOrganizationMemberResponse) {
	response = &CreateOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集团成员账号
func (c *Client) CreateOrganizationMember(request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationMemberRequest()
	}
	response = NewCreateOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationMembersRequest() (request *DeleteOrganizationMembersRequest) {
	request = &DeleteOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembers")
	return
}

func NewDeleteOrganizationMembersResponse() (response *DeleteOrganizationMembersResponse) {
	response = &DeleteOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除成员账号
func (c *Client) DeleteOrganizationMembers(request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationMembersRequest()
	}
	response = NewDeleteOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationRequest() (request *UpdateOrganizationRequest) {
	request = &UpdateOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganization")
	return
}

func NewUpdateOrganizationResponse() (response *UpdateOrganizationResponse) {
	response = &UpdateOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新集团组织信息
func (c *Client) UpdateOrganization(request *UpdateOrganizationRequest) (response *UpdateOrganizationResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationRequest()
	}
	response = NewUpdateOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationMemberRequest() (request *UpdateOrganizationMemberRequest) {
	request = &UpdateOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMember")
	return
}

func NewUpdateOrganizationMemberResponse() (response *UpdateOrganizationMemberResponse) {
	response = &UpdateOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新成员账号
func (c *Client) UpdateOrganizationMember(request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationMemberRequest()
	}
	response = NewUpdateOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
	request = &DescribeOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembers")
	return
}

func NewDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
	response = &DescribeOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 成员账号列表
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMembersRequest()
	}
	response = NewDescribeOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationRequest() (request *CreateOrganizationRequest) {
	request = &CreateOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganization")
	return
}

func NewCreateOrganizationResponse() (response *CreateOrganizationResponse) {
	response = &CreateOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集团组织
func (c *Client) CreateOrganization(request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationRequest()
	}
	response = NewCreateOrganizationResponse()
	err = c.Send(request, response)
	return
}
