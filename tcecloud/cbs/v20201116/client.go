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

package v20201116

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-11-16"

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

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
	request = &DescribeSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeSnapshots")
	return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
	response = &DescribeSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeSnapshots）用于查询快照的详细信息。
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
	if request == nil {
		request = NewDescribeSnapshotsRequest()
	}
	response = NewDescribeSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDiskAssociatedSnapshotsRequest() (request *DescribeDiskAssociatedSnapshotsRequest) {
	request = &DescribeDiskAssociatedSnapshotsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDiskAssociatedSnapshots")
	return
}

func NewDescribeDiskAssociatedSnapshotsResponse() (response *DescribeDiskAssociatedSnapshotsResponse) {
	response = &DescribeDiskAssociatedSnapshotsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询云硬盘关联的所有快照信息
func (c *Client) DescribeDiskAssociatedSnapshots(request *DescribeDiskAssociatedSnapshotsRequest) (response *DescribeDiskAssociatedSnapshotsResponse, err error) {
	if request == nil {
		request = NewDescribeDiskAssociatedSnapshotsRequest()
	}
	response = NewDescribeDiskAssociatedSnapshotsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cbs", APIVersion, "DescribeDisks")
	return
}

func NewDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeDisks）用于查询云硬盘列表。
//
// * 可以根据云硬盘ID、云硬盘类型或者云硬盘状态等信息来查询云硬盘的详细信息，不同条件之间为与(AND)的关系，过滤信息详细请见过滤器`Filter`。
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的云硬盘列表。
func (c *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	if request == nil {
		request = NewDescribeDisksRequest()
	}
	response = NewDescribeDisksResponse()
	err = c.Send(request, response)
	return
}
