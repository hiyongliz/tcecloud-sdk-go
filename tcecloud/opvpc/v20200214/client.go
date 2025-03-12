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

package v20200214

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-02-14"

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

func NewGetNfvGwFormatsRequest() (request *GetNfvGwFormatsRequest) {
	request = &GetNfvGwFormatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvGwFormats")
	return
}

func NewGetNfvGwFormatsResponse() (response *GetNfvGwFormatsResponse) {
	response = &GetNfvGwFormatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Nfv网关的规格
func (c *Client) GetNfvGwFormats(request *GetNfvGwFormatsRequest) (response *GetNfvGwFormatsResponse, err error) {
	if request == nil {
		request = NewGetNfvGwFormatsRequest()
	}
	response = NewGetNfvGwFormatsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayIpsRequest() (request *DescribeGatewayIpsRequest) {
	request = &DescribeGatewayIpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayIps")
	return
}

func NewDescribeGatewayIpsResponse() (response *DescribeGatewayIpsResponse) {
	response = &DescribeGatewayIpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关ip
func (c *Client) DescribeGatewayIps(request *DescribeGatewayIpsRequest) (response *DescribeGatewayIpsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayIpsRequest()
	}
	response = NewDescribeGatewayIpsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointServiceRequest() (request *DescribeVpcEndPointServiceRequest) {
	request = &DescribeVpcEndPointServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcEndPointService")
	return
}

func NewDescribeVpcEndPointServiceResponse() (response *DescribeVpcEndPointServiceResponse) {
	response = &DescribeVpcEndPointServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取终端节点列表。
func (c *Client) DescribeVpcEndPointService(request *DescribeVpcEndPointServiceRequest) (response *DescribeVpcEndPointServiceResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointServiceRequest()
	}
	response = NewDescribeVpcEndPointServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUnderlayNatGatewayRuleRequest() (request *DeleteUnderlayNatGatewayRuleRequest) {
	request = &DeleteUnderlayNatGatewayRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteUnderlayNatGatewayRule")
	return
}

func NewDeleteUnderlayNatGatewayRuleResponse() (response *DeleteUnderlayNatGatewayRuleResponse) {
	response = &DeleteUnderlayNatGatewayRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除underlay natgateway规则
func (c *Client) DeleteUnderlayNatGatewayRule(request *DeleteUnderlayNatGatewayRuleRequest) (response *DeleteUnderlayNatGatewayRuleResponse, err error) {
	if request == nil {
		request = NewDeleteUnderlayNatGatewayRuleRequest()
	}
	response = NewDeleteUnderlayNatGatewayRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAclListRequest() (request *DescribeAclListRequest) {
	request = &DescribeAclListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeAclList")
	return
}

func NewDescribeAclListResponse() (response *DescribeAclListResponse) {
	response = &DescribeAclListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取acl列表
func (c *Client) DescribeAclList(request *DescribeAclListRequest) (response *DescribeAclListResponse, err error) {
	if request == nil {
		request = NewDescribeAclListRequest()
	}
	response = NewDescribeAclListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVsrsRequest() (request *DescribeVsrsRequest) {
	request = &DescribeVsrsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVsrs")
	return
}

func NewDescribeVsrsResponse() (response *DescribeVsrsResponse) {
	response = &DescribeVsrsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VSR列表
func (c *Client) DescribeVsrs(request *DescribeVsrsRequest) (response *DescribeVsrsResponse, err error) {
	if request == nil {
		request = NewDescribeVsrsRequest()
	}
	response = NewDescribeVsrsResponse()
	err = c.Send(request, response)
	return
}

func NewGetNfvGwListRequest() (request *GetNfvGwListRequest) {
	request = &GetNfvGwListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvGwList")
	return
}

func NewGetNfvGwListResponse() (response *GetNfvGwListResponse) {
	response = &GetNfvGwListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Nfv网关的列表
func (c *Client) GetNfvGwList(request *GetNfvGwListRequest) (response *GetNfvGwListResponse, err error) {
	if request == nil {
		request = NewGetNfvGwListRequest()
	}
	response = NewGetNfvGwListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubnetsExRequest() (request *CreateSubnetsExRequest) {
	request = &CreateSubnetsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateSubnetsEx")
	return
}

func NewCreateSubnetsExResponse() (response *CreateSubnetsExResponse) {
	response = &CreateSubnetsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创建子网
func (c *Client) CreateSubnetsEx(request *CreateSubnetsExRequest) (response *CreateSubnetsExResponse, err error) {
	if request == nil {
		request = NewCreateSubnetsExRequest()
	}
	response = NewCreateSubnetsExResponse()
	err = c.Send(request, response)
	return
}

func NewGetNfvGwLogsRequest() (request *GetNfvGwLogsRequest) {
	request = &GetNfvGwLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvGwLogs")
	return
}

func NewGetNfvGwLogsResponse() (response *GetNfvGwLogsResponse) {
	response = &GetNfvGwLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Nfv网关的日志
func (c *Client) GetNfvGwLogs(request *GetNfvGwLogsRequest) (response *GetNfvGwLogsResponse, err error) {
	if request == nil {
		request = NewGetNfvGwLogsRequest()
	}
	response = NewGetNfvGwLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnConnsRequest() (request *DescribeVpnConnsRequest) {
	request = &DescribeVpnConnsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpnConns")
	return
}

func NewDescribeVpnConnsResponse() (response *DescribeVpnConnsResponse) {
	response = &DescribeVpnConnsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vpn通道列表
func (c *Client) DescribeVpnConns(request *DescribeVpnConnsRequest) (response *DescribeVpnConnsResponse, err error) {
	if request == nil {
		request = NewDescribeVpnConnsRequest()
	}
	response = NewDescribeVpnConnsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBasicNetsRequest() (request *DescribeBasicNetsRequest) {
	request = &DescribeBasicNetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeBasicNets")
	return
}

func NewDescribeBasicNetsResponse() (response *DescribeBasicNetsResponse) {
	response = &DescribeBasicNetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPC0网段列表
func (c *Client) DescribeBasicNets(request *DescribeBasicNetsRequest) (response *DescribeBasicNetsResponse, err error) {
	if request == nil {
		request = NewDescribeBasicNetsRequest()
	}
	response = NewDescribeBasicNetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayWanInfoRequest() (request *DescribeUnderlayNatGatewayWanInfoRequest) {
	request = &DescribeUnderlayNatGatewayWanInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayWanInfo")
	return
}

func NewDescribeUnderlayNatGatewayWanInfoResponse() (response *DescribeUnderlayNatGatewayWanInfoResponse) {
	response = &DescribeUnderlayNatGatewayWanInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取underlay natgateway运营商信息
func (c *Client) DescribeUnderlayNatGatewayWanInfo(request *DescribeUnderlayNatGatewayWanInfoRequest) (response *DescribeUnderlayNatGatewayWanInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayWanInfoRequest()
	}
	response = NewDescribeUnderlayNatGatewayWanInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayCLBInstallationRequest() (request *CreateGatewayCLBInstallationRequest) {
	request = &CreateGatewayCLBInstallationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateGatewayCLBInstallation")
	return
}

func NewCreateGatewayCLBInstallationResponse() (response *CreateGatewayCLBInstallationResponse) {
	response = &CreateGatewayCLBInstallationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateGatewayCLBInstallation
func (c *Client) CreateGatewayCLBInstallation(request *CreateGatewayCLBInstallationRequest) (response *CreateGatewayCLBInstallationResponse, err error) {
	if request == nil {
		request = NewCreateGatewayCLBInstallationRequest()
	}
	response = NewCreateGatewayCLBInstallationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUnderlayNatGatewayRuleWipRequest() (request *DeleteUnderlayNatGatewayRuleWipRequest) {
	request = &DeleteUnderlayNatGatewayRuleWipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteUnderlayNatGatewayRuleWip")
	return
}

func NewDeleteUnderlayNatGatewayRuleWipResponse() (response *DeleteUnderlayNatGatewayRuleWipResponse) {
	response = &DeleteUnderlayNatGatewayRuleWipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除underlay natgateway规则运营商外网ip
func (c *Client) DeleteUnderlayNatGatewayRuleWip(request *DeleteUnderlayNatGatewayRuleWipRequest) (response *DeleteUnderlayNatGatewayRuleWipResponse, err error) {
	if request == nil {
		request = NewDeleteUnderlayNatGatewayRuleWipRequest()
	}
	response = NewDeleteUnderlayNatGatewayRuleWipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayWanIpsRequest() (request *DescribeUnderlayNatGatewayWanIpsRequest) {
	request = &DescribeUnderlayNatGatewayWanIpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayWanIps")
	return
}

func NewDescribeUnderlayNatGatewayWanIpsResponse() (response *DescribeUnderlayNatGatewayWanIpsResponse) {
	response = &DescribeUnderlayNatGatewayWanIpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取underlay natgateway运营商ip
func (c *Client) DescribeUnderlayNatGatewayWanIps(request *DescribeUnderlayNatGatewayWanIpsRequest) (response *DescribeUnderlayNatGatewayWanIpsResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayWanIpsRequest()
	}
	response = NewDescribeUnderlayNatGatewayWanIpsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVmListRequest() (request *DescribeVmListRequest) {
	request = &DescribeVmListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVmList")
	return
}

func NewDescribeVmListResponse() (response *DescribeVmListResponse) {
	response = &DescribeVmListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子机列表
func (c *Client) DescribeVmList(request *DescribeVmListRequest) (response *DescribeVmListResponse, err error) {
	if request == nil {
		request = NewDescribeVmListRequest()
	}
	response = NewDescribeVmListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDetectRouteStateRequest() (request *ModifyDetectRouteStateRequest) {
	request = &ModifyDetectRouteStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyDetectRouteState")
	return
}

func NewModifyDetectRouteStateResponse() (response *ModifyDetectRouteStateResponse) {
	response = &ModifyDetectRouteStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 探测机任务重置。
func (c *Client) ModifyDetectRouteState(request *ModifyDetectRouteStateRequest) (response *ModifyDetectRouteStateResponse, err error) {
	if request == nil {
		request = NewModifyDetectRouteStateRequest()
	}
	response = NewModifyDetectRouteStateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServiceAttributeRequest() (request *ModifyServiceAttributeRequest) {
	request = &ModifyServiceAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyServiceAttribute")
	return
}

func NewModifyServiceAttributeResponse() (response *ModifyServiceAttributeResponse) {
	response = &ModifyServiceAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改物理网络服务
func (c *Client) ModifyServiceAttribute(request *ModifyServiceAttributeRequest) (response *ModifyServiceAttributeResponse, err error) {
	if request == nil {
		request = NewModifyServiceAttributeRequest()
	}
	response = NewModifyServiceAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupForCvmRequest() (request *DescribeSecurityGroupForCvmRequest) {
	request = &DescribeSecurityGroupForCvmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSecurityGroupForCvm")
	return
}

func NewDescribeSecurityGroupForCvmResponse() (response *DescribeSecurityGroupForCvmResponse) {
	response = &DescribeSecurityGroupForCvmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vm关联的安全组
func (c *Client) DescribeSecurityGroupForCvm(request *DescribeSecurityGroupForCvmRequest) (response *DescribeSecurityGroupForCvmResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupForCvmRequest()
	}
	response = NewDescribeSecurityGroupForCvmResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateServiceRouteRequest() (request *MigrateServiceRouteRequest) {
	request = &MigrateServiceRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "MigrateServiceRoute")
	return
}

func NewMigrateServiceRouteResponse() (response *MigrateServiceRouteResponse) {
	response = &MigrateServiceRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移物理网络服务映射规则。
func (c *Client) MigrateServiceRoute(request *MigrateServiceRouteRequest) (response *MigrateServiceRouteResponse, err error) {
	if request == nil {
		request = NewMigrateServiceRouteRequest()
	}
	response = NewMigrateServiceRouteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUnderlayNatGatewayWanInfoAndWanIpRequest() (request *CreateUnderlayNatGatewayWanInfoAndWanIpRequest) {
	request = &CreateUnderlayNatGatewayWanInfoAndWanIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateUnderlayNatGatewayWanInfoAndWanIp")
	return
}

func NewCreateUnderlayNatGatewayWanInfoAndWanIpResponse() (response *CreateUnderlayNatGatewayWanInfoAndWanIpResponse) {
	response = &CreateUnderlayNatGatewayWanInfoAndWanIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建underlay natgateway运营商信息和对应的外网ip
func (c *Client) CreateUnderlayNatGatewayWanInfoAndWanIp(request *CreateUnderlayNatGatewayWanInfoAndWanIpRequest) (response *CreateUnderlayNatGatewayWanInfoAndWanIpResponse, err error) {
	if request == nil {
		request = NewCreateUnderlayNatGatewayWanInfoAndWanIpRequest()
	}
	response = NewCreateUnderlayNatGatewayWanInfoAndWanIpResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubnetsRequest() (request *DeleteSubnetsRequest) {
	request = &DeleteSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteSubnets")
	return
}

func NewDeleteSubnetsResponse() (response *DeleteSubnetsResponse) {
	response = &DeleteSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除子网
func (c *Client) DeleteSubnets(request *DeleteSubnetsRequest) (response *DeleteSubnetsResponse, err error) {
	if request == nil {
		request = NewDeleteSubnetsRequest()
	}
	response = NewDeleteSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateGatewayRequest() (request *MigrateGatewayRequest) {
	request = &MigrateGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "MigrateGateway")
	return
}

func NewMigrateGatewayResponse() (response *MigrateGatewayResponse) {
	response = &MigrateGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 迁移网关
func (c *Client) MigrateGateway(request *MigrateGatewayRequest) (response *MigrateGatewayResponse, err error) {
	if request == nil {
		request = NewMigrateGatewayRequest()
	}
	response = NewMigrateGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUnderlayNatGatewayWanIpRequest() (request *DeleteUnderlayNatGatewayWanIpRequest) {
	request = &DeleteUnderlayNatGatewayWanIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteUnderlayNatGatewayWanIp")
	return
}

func NewDeleteUnderlayNatGatewayWanIpResponse() (response *DeleteUnderlayNatGatewayWanIpResponse) {
	response = &DeleteUnderlayNatGatewayWanIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除underlay natgateway运营商ip
func (c *Client) DeleteUnderlayNatGatewayWanIp(request *DeleteUnderlayNatGatewayWanIpRequest) (response *DeleteUnderlayNatGatewayWanIpResponse, err error) {
	if request == nil {
		request = NewDeleteUnderlayNatGatewayWanIpRequest()
	}
	response = NewDeleteUnderlayNatGatewayWanIpResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayGroupsRequest() (request *DescribeGatewayGroupsRequest) {
	request = &DescribeGatewayGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayGroups")
	return
}

func NewDescribeGatewayGroupsResponse() (response *DescribeGatewayGroupsResponse) {
	response = &DescribeGatewayGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关集群组列表
func (c *Client) DescribeGatewayGroups(request *DescribeGatewayGroupsRequest) (response *DescribeGatewayGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayGroupsRequest()
	}
	response = NewDescribeGatewayGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnGatewayDetailRequest() (request *DescribeVpnGatewayDetailRequest) {
	request = &DescribeVpnGatewayDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpnGatewayDetail")
	return
}

func NewDescribeVpnGatewayDetailResponse() (response *DescribeVpnGatewayDetailResponse) {
	response = &DescribeVpnGatewayDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vpn本地网关详情
func (c *Client) DescribeVpnGatewayDetail(request *DescribeVpnGatewayDetailRequest) (response *DescribeVpnGatewayDetailResponse, err error) {
	if request == nil {
		request = NewDescribeVpnGatewayDetailRequest()
	}
	response = NewDescribeVpnGatewayDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcGlobalExtendCidrRequest() (request *ModifyVpcGlobalExtendCidrRequest) {
	request = &ModifyVpcGlobalExtendCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyVpcGlobalExtendCidr")
	return
}

func NewModifyVpcGlobalExtendCidrResponse() (response *ModifyVpcGlobalExtendCidrResponse) {
	response = &ModifyVpcGlobalExtendCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyVpcGlobalExtendCidr
func (c *Client) ModifyVpcGlobalExtendCidr(request *ModifyVpcGlobalExtendCidrRequest) (response *ModifyVpcGlobalExtendCidrResponse, err error) {
	if request == nil {
		request = NewModifyVpcGlobalExtendCidrRequest()
	}
	response = NewModifyVpcGlobalExtendCidrResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcExRequest() (request *DescribeVpcExRequest) {
	request = &DescribeVpcExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcEx")
	return
}

func NewDescribeVpcExResponse() (response *DescribeVpcExResponse) {
	response = &DescribeVpcExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vpc列表
func (c *Client) DescribeVpcEx(request *DescribeVpcExRequest) (response *DescribeVpcExResponse, err error) {
	if request == nil {
		request = NewDescribeVpcExRequest()
	}
	response = NewDescribeVpcExResponse()
	err = c.Send(request, response)
	return
}

func NewMigrateDirectConnectGatewayRequest() (request *MigrateDirectConnectGatewayRequest) {
	request = &MigrateDirectConnectGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "MigrateDirectConnectGateway")
	return
}

func NewMigrateDirectConnectGatewayResponse() (response *MigrateDirectConnectGatewayResponse) {
	response = &MigrateDirectConnectGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专线网关迁移
func (c *Client) MigrateDirectConnectGateway(request *MigrateDirectConnectGatewayRequest) (response *MigrateDirectConnectGatewayResponse, err error) {
	if request == nil {
		request = NewMigrateDirectConnectGatewayRequest()
	}
	response = NewMigrateDirectConnectGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayTypeNamesRequest() (request *DescribeGatewayTypeNamesRequest) {
	request = &DescribeGatewayTypeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayTypeNames")
	return
}

func NewDescribeGatewayTypeNamesResponse() (response *DescribeGatewayTypeNamesResponse) {
	response = &DescribeGatewayTypeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关类型名字列表
func (c *Client) DescribeGatewayTypeNames(request *DescribeGatewayTypeNamesRequest) (response *DescribeGatewayTypeNamesResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayTypeNamesRequest()
	}
	response = NewDescribeGatewayTypeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewGetNfvGwDetailRequest() (request *GetNfvGwDetailRequest) {
	request = &GetNfvGwDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvGwDetail")
	return
}

func NewGetNfvGwDetailResponse() (response *GetNfvGwDetailResponse) {
	response = &GetNfvGwDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetNfvGwDetail
func (c *Client) GetNfvGwDetail(request *GetNfvGwDetailRequest) (response *GetNfvGwDetailResponse, err error) {
	if request == nil {
		request = NewGetNfvGwDetailRequest()
	}
	response = NewGetNfvGwDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNfvGwNodeRequest() (request *DeleteNfvGwNodeRequest) {
	request = &DeleteNfvGwNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteNfvGwNode")
	return
}

func NewDeleteNfvGwNodeResponse() (response *DeleteNfvGwNodeResponse) {
	response = &DeleteNfvGwNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Nfv网关节点
func (c *Client) DeleteNfvGwNode(request *DeleteNfvGwNodeRequest) (response *DeleteNfvGwNodeResponse, err error) {
	if request == nil {
		request = NewDeleteNfvGwNodeRequest()
	}
	response = NewDeleteNfvGwNodeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVpcLimitRequest() (request *ModifyVpcLimitRequest) {
	request = &ModifyVpcLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyVpcLimit")
	return
}

func NewModifyVpcLimitResponse() (response *ModifyVpcLimitResponse) {
	response = &ModifyVpcLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改vpc配额信息
func (c *Client) ModifyVpcLimit(request *ModifyVpcLimitRequest) (response *ModifyVpcLimitResponse, err error) {
	if request == nil {
		request = NewModifyVpcLimitRequest()
	}
	response = NewModifyVpcLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUnderlayNatGatewayRuleRipRequest() (request *DeleteUnderlayNatGatewayRuleRipRequest) {
	request = &DeleteUnderlayNatGatewayRuleRipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteUnderlayNatGatewayRuleRip")
	return
}

func NewDeleteUnderlayNatGatewayRuleRipResponse() (response *DeleteUnderlayNatGatewayRuleRipResponse) {
	response = &DeleteUnderlayNatGatewayRuleRipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除underlay natgateway规则对应rip
func (c *Client) DeleteUnderlayNatGatewayRuleRip(request *DeleteUnderlayNatGatewayRuleRipRequest) (response *DeleteUnderlayNatGatewayRuleRipResponse, err error) {
	if request == nil {
		request = NewDeleteUnderlayNatGatewayRuleRipRequest()
	}
	response = NewDeleteUnderlayNatGatewayRuleRipResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductFailureMigrateTaskDetailRequest() (request *QueryProductFailureMigrateTaskDetailRequest) {
	request = &QueryProductFailureMigrateTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "QueryProductFailureMigrateTaskDetail")
	return
}

func NewQueryProductFailureMigrateTaskDetailResponse() (response *QueryProductFailureMigrateTaskDetailResponse) {
	response = &QueryProductFailureMigrateTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品故障切换结果查询
func (c *Client) QueryProductFailureMigrateTaskDetail(request *QueryProductFailureMigrateTaskDetailRequest) (response *QueryProductFailureMigrateTaskDetailResponse, err error) {
	if request == nil {
		request = NewQueryProductFailureMigrateTaskDetailRequest()
	}
	response = NewQueryProductFailureMigrateTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetDetectNodeInfoRequest() (request *GetDetectNodeInfoRequest) {
	request = &GetDetectNodeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetDetectNodeInfo")
	return
}

func NewGetDetectNodeInfoResponse() (response *GetDetectNodeInfoResponse) {
	response = &GetDetectNodeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取探测机节点信息
func (c *Client) GetDetectNodeInfo(request *GetDetectNodeInfoRequest) (response *GetDetectNodeInfoResponse, err error) {
	if request == nil {
		request = NewGetDetectNodeInfoRequest()
	}
	response = NewGetDetectNodeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayCarrierInfoRequest() (request *DescribeUnderlayNatGatewayCarrierInfoRequest) {
	request = &DescribeUnderlayNatGatewayCarrierInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayCarrierInfo")
	return
}

func NewDescribeUnderlayNatGatewayCarrierInfoResponse() (response *DescribeUnderlayNatGatewayCarrierInfoResponse) {
	response = &DescribeUnderlayNatGatewayCarrierInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公网运营商列表
func (c *Client) DescribeUnderlayNatGatewayCarrierInfo(request *DescribeUnderlayNatGatewayCarrierInfoRequest) (response *DescribeUnderlayNatGatewayCarrierInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayCarrierInfoRequest()
	}
	response = NewDescribeUnderlayNatGatewayCarrierInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDetectHostRequest() (request *DeleteDetectHostRequest) {
	request = &DeleteDetectHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteDetectHost")
	return
}

func NewDeleteDetectHostResponse() (response *DeleteDetectHostResponse) {
	response = &DeleteDetectHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除探测机
func (c *Client) DeleteDetectHost(request *DeleteDetectHostRequest) (response *DeleteDetectHostResponse, err error) {
	if request == nil {
		request = NewDeleteDetectHostRequest()
	}
	response = NewDeleteDetectHostResponse()
	err = c.Send(request, response)
	return
}

func NewGetSwitchGroupRequest() (request *GetSwitchGroupRequest) {
	request = &GetSwitchGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetSwitchGroup")
	return
}

func NewGetSwitchGroupResponse() (response *GetSwitchGroupResponse) {
	response = &GetSwitchGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取交换机组信息
func (c *Client) GetSwitchGroup(request *GetSwitchGroupRequest) (response *GetSwitchGroupResponse, err error) {
	if request == nil {
		request = NewGetSwitchGroupRequest()
	}
	response = NewGetSwitchGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteJnsGatewayServiceRequest() (request *DeleteJnsGatewayServiceRequest) {
	request = &DeleteJnsGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteJnsGatewayService")
	return
}

func NewDeleteJnsGatewayServiceResponse() (response *DeleteJnsGatewayServiceResponse) {
	response = &DeleteJnsGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除私有网络映射服务
func (c *Client) DeleteJnsGatewayService(request *DeleteJnsGatewayServiceRequest) (response *DeleteJnsGatewayServiceResponse, err error) {
	if request == nil {
		request = NewDeleteJnsGatewayServiceRequest()
	}
	response = NewDeleteJnsGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceRequest() (request *DeleteServiceRequest) {
	request = &DeleteServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteService")
	return
}

func NewDeleteServiceResponse() (response *DeleteServiceResponse) {
	response = &DeleteServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除物理网络服务
func (c *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
	if request == nil {
		request = NewDeleteServiceRequest()
	}
	response = NewDeleteServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewayRequest() (request *DeleteGatewayRequest) {
	request = &DeleteGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteGateway")
	return
}

func NewDeleteGatewayResponse() (response *DeleteGatewayResponse) {
	response = &DeleteGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关集群
func (c *Client) DeleteGateway(request *DeleteGatewayRequest) (response *DeleteGatewayResponse, err error) {
	if request == nil {
		request = NewDeleteGatewayRequest()
	}
	response = NewDeleteGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDetectPkgVersionsRequest() (request *DescribeDetectPkgVersionsRequest) {
	request = &DescribeDetectPkgVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeDetectPkgVersions")
	return
}

func NewDescribeDetectPkgVersionsResponse() (response *DescribeDetectPkgVersionsResponse) {
	response = &DescribeDetectPkgVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取探测机安装包版本列表
func (c *Client) DescribeDetectPkgVersions(request *DescribeDetectPkgVersionsRequest) (response *DescribeDetectPkgVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeDetectPkgVersionsRequest()
	}
	response = NewDescribeDetectPkgVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductHealthStateRequest() (request *QueryProductHealthStateRequest) {
	request = &QueryProductHealthStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "QueryProductHealthState")
	return
}

func NewQueryProductHealthStateResponse() (response *QueryProductHealthStateResponse) {
	response = &QueryProductHealthStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询VPC产品健康状态，结果中应该包盖该产品一个或者多个维度的健康状态和总体健康状态。
func (c *Client) QueryProductHealthState(request *QueryProductHealthStateRequest) (response *QueryProductHealthStateResponse, err error) {
	if request == nil {
		request = NewQueryProductHealthStateRequest()
	}
	response = NewQueryProductHealthStateResponse()
	err = c.Send(request, response)
	return
}

func NewAddUnderlayNatGatewayRuleWipRequest() (request *AddUnderlayNatGatewayRuleWipRequest) {
	request = &AddUnderlayNatGatewayRuleWipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "AddUnderlayNatGatewayRuleWip")
	return
}

func NewAddUnderlayNatGatewayRuleWipResponse() (response *AddUnderlayNatGatewayRuleWipResponse) {
	response = &AddUnderlayNatGatewayRuleWipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增underlay natgateway规则运营商外网ip
func (c *Client) AddUnderlayNatGatewayRuleWip(request *AddUnderlayNatGatewayRuleWipRequest) (response *AddUnderlayNatGatewayRuleWipResponse, err error) {
	if request == nil {
		request = NewAddUnderlayNatGatewayRuleWipRequest()
	}
	response = NewAddUnderlayNatGatewayRuleWipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDetectInstallStateRequest() (request *DescribeDetectInstallStateRequest) {
	request = &DescribeDetectInstallStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeDetectInstallState")
	return
}

func NewDescribeDetectInstallStateResponse() (response *DescribeDetectInstallStateResponse) {
	response = &DescribeDetectInstallStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取探测机自动化部署状态
func (c *Client) DescribeDetectInstallState(request *DescribeDetectInstallStateRequest) (response *DescribeDetectInstallStateResponse, err error) {
	if request == nil {
		request = NewDescribeDetectInstallStateRequest()
	}
	response = NewDescribeDetectInstallStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpnGatewaysRequest() (request *DescribeVpnGatewaysRequest) {
	request = &DescribeVpnGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpnGateways")
	return
}

func NewDescribeVpnGatewaysResponse() (response *DescribeVpnGatewaysResponse) {
	response = &DescribeVpnGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vpn本地网关列表
func (c *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeVpnGatewaysRequest()
	}
	response = NewDescribeVpnGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcPeersRequest() (request *DescribeVpcPeersRequest) {
	request = &DescribeVpcPeersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcPeers")
	return
}

func NewDescribeVpcPeersResponse() (response *DescribeVpcPeersResponse) {
	response = &DescribeVpcPeersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询对等连接列表
func (c *Client) DescribeVpcPeers(request *DescribeVpcPeersRequest) (response *DescribeVpcPeersResponse, err error) {
	if request == nil {
		request = NewDescribeVpcPeersRequest()
	}
	response = NewDescribeVpcPeersResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackNfvGwLogStepRequest() (request *RollbackNfvGwLogStepRequest) {
	request = &RollbackNfvGwLogStepRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "RollbackNfvGwLogStep")
	return
}

func NewRollbackNfvGwLogStepResponse() (response *RollbackNfvGwLogStepResponse) {
	response = &RollbackNfvGwLogStepResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚启动Nfv网关日志步骤
func (c *Client) RollbackNfvGwLogStep(request *RollbackNfvGwLogStepRequest) (response *RollbackNfvGwLogStepResponse, err error) {
	if request == nil {
		request = NewRollbackNfvGwLogStepRequest()
	}
	response = NewRollbackNfvGwLogStepResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNfvGwRequest() (request *DeleteNfvGwRequest) {
	request = &DeleteNfvGwRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteNfvGw")
	return
}

func NewDeleteNfvGwResponse() (response *DeleteNfvGwResponse) {
	response = &DeleteNfvGwResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Nfv网关
func (c *Client) DeleteNfvGw(request *DeleteNfvGwRequest) (response *DeleteNfvGwResponse, err error) {
	if request == nil {
		request = NewDeleteNfvGwRequest()
	}
	response = NewDeleteNfvGwResponse()
	err = c.Send(request, response)
	return
}

func NewGetClbGatewayIfcfgRequest() (request *GetClbGatewayIfcfgRequest) {
	request = &GetClbGatewayIfcfgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetClbGatewayIfcfg")
	return
}

func NewGetClbGatewayIfcfgResponse() (response *GetClbGatewayIfcfgResponse) {
	response = &GetClbGatewayIfcfgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关lan_segment数据信息
func (c *Client) GetClbGatewayIfcfg(request *GetClbGatewayIfcfgRequest) (response *GetClbGatewayIfcfgResponse, err error) {
	if request == nil {
		request = NewGetClbGatewayIfcfgRequest()
	}
	response = NewGetClbGatewayIfcfgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVmIpv6InfoRequest() (request *DescribeVmIpv6InfoRequest) {
	request = &DescribeVmIpv6InfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVmIpv6Info")
	return
}

func NewDescribeVmIpv6InfoResponse() (response *DescribeVmIpv6InfoResponse) {
	response = &DescribeVmIpv6InfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子机ipv6信息
func (c *Client) DescribeVmIpv6Info(request *DescribeVmIpv6InfoRequest) (response *DescribeVmIpv6InfoResponse, err error) {
	if request == nil {
		request = NewDescribeVmIpv6InfoRequest()
	}
	response = NewDescribeVmIpv6InfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayClbInstallStateRequest() (request *DescribeGatewayClbInstallStateRequest) {
	request = &DescribeGatewayClbInstallStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayClbInstallState")
	return
}

func NewDescribeGatewayClbInstallStateResponse() (response *DescribeGatewayClbInstallStateResponse) {
	response = &DescribeGatewayClbInstallStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取clb网关自动化部署结果
func (c *Client) DescribeGatewayClbInstallState(request *DescribeGatewayClbInstallStateRequest) (response *DescribeGatewayClbInstallStateResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayClbInstallStateRequest()
	}
	response = NewDescribeGatewayClbInstallStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayRuleRipsRequest() (request *DescribeUnderlayNatGatewayRuleRipsRequest) {
	request = &DescribeUnderlayNatGatewayRuleRipsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayRuleRips")
	return
}

func NewDescribeUnderlayNatGatewayRuleRipsResponse() (response *DescribeUnderlayNatGatewayRuleRipsResponse) {
	response = &DescribeUnderlayNatGatewayRuleRipsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取underlay natgateway规则对应运营端rip
func (c *Client) DescribeUnderlayNatGatewayRuleRips(request *DescribeUnderlayNatGatewayRuleRipsRequest) (response *DescribeUnderlayNatGatewayRuleRipsResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayRuleRipsRequest()
	}
	response = NewDescribeUnderlayNatGatewayRuleRipsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClbGatewayVersionsRequest() (request *DescribeClbGatewayVersionsRequest) {
	request = &DescribeClbGatewayVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeClbGatewayVersions")
	return
}

func NewDescribeClbGatewayVersionsResponse() (response *DescribeClbGatewayVersionsResponse) {
	response = &DescribeClbGatewayVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取clb网关版本信息
func (c *Client) DescribeClbGatewayVersions(request *DescribeClbGatewayVersionsRequest) (response *DescribeClbGatewayVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeClbGatewayVersionsRequest()
	}
	response = NewDescribeClbGatewayVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
	request = &DescribeNatGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeNatGateways")
	return
}

func NewDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
	response = &DescribeNatGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取nat 网关列表
func (c *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeNatGatewaysRequest()
	}
	response = NewDescribeNatGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtectPolicyIpRequest() (request *DescribeProtectPolicyIpRequest) {
	request = &DescribeProtectPolicyIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeProtectPolicyIp")
	return
}

func NewDescribeProtectPolicyIpResponse() (response *DescribeProtectPolicyIpResponse) {
	response = &DescribeProtectPolicyIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关过载保护信息。
func (c *Client) DescribeProtectPolicyIp(request *DescribeProtectPolicyIpRequest) (response *DescribeProtectPolicyIpResponse, err error) {
	if request == nil {
		request = NewDescribeProtectPolicyIpRequest()
	}
	response = NewDescribeProtectPolicyIpResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIpV6SubnetsRequest() (request *CreateIpV6SubnetsRequest) {
	request = &CreateIpV6SubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateIpV6Subnets")
	return
}

func NewCreateIpV6SubnetsResponse() (response *CreateIpV6SubnetsResponse) {
	response = &CreateIpV6SubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 投放VPC IPv6地址段
func (c *Client) CreateIpV6Subnets(request *CreateIpV6SubnetsRequest) (response *CreateIpV6SubnetsResponse, err error) {
	if request == nil {
		request = NewCreateIpV6SubnetsRequest()
	}
	response = NewCreateIpV6SubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchUnderlayNatGatewayHaStatusRequest() (request *SwitchUnderlayNatGatewayHaStatusRequest) {
	request = &SwitchUnderlayNatGatewayHaStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "SwitchUnderlayNatGatewayHaStatus")
	return
}

func NewSwitchUnderlayNatGatewayHaStatusResponse() (response *SwitchUnderlayNatGatewayHaStatusResponse) {
	response = &SwitchUnderlayNatGatewayHaStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SwitchUnderlayNatGatewayHaStatus
func (c *Client) SwitchUnderlayNatGatewayHaStatus(request *SwitchUnderlayNatGatewayHaStatusRequest) (response *SwitchUnderlayNatGatewayHaStatusResponse, err error) {
	if request == nil {
		request = NewSwitchUnderlayNatGatewayHaStatusRequest()
	}
	response = NewSwitchUnderlayNatGatewayHaStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVpcPeerCrossRegionRequest() (request *DeleteVpcPeerCrossRegionRequest) {
	request = &DeleteVpcPeerCrossRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteVpcPeerCrossRegion")
	return
}

func NewDeleteVpcPeerCrossRegionResponse() (response *DeleteVpcPeerCrossRegionResponse) {
	response = &DeleteVpcPeerCrossRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除对等连接有效的地域线路
func (c *Client) DeleteVpcPeerCrossRegion(request *DeleteVpcPeerCrossRegionRequest) (response *DeleteVpcPeerCrossRegionResponse, err error) {
	if request == nil {
		request = NewDeleteVpcPeerCrossRegionRequest()
	}
	response = NewDeleteVpcPeerCrossRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHaVipListRequest() (request *DescribeHaVipListRequest) {
	request = &DescribeHaVipListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeHaVipList")
	return
}

func NewDescribeHaVipListResponse() (response *DescribeHaVipListResponse) {
	response = &DescribeHaVipListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端查询HAVIP列表
func (c *Client) DescribeHaVipList(request *DescribeHaVipListRequest) (response *DescribeHaVipListResponse, err error) {
	if request == nil {
		request = NewDescribeHaVipListRequest()
	}
	response = NewDescribeHaVipListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcLimitNamesToTypeRequest() (request *DescribeVpcLimitNamesToTypeRequest) {
	request = &DescribeVpcLimitNamesToTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcLimitNamesToType")
	return
}

func NewDescribeVpcLimitNamesToTypeResponse() (response *DescribeVpcLimitNamesToTypeResponse) {
	response = &DescribeVpcLimitNamesToTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配额管理名称和类型映射列表
func (c *Client) DescribeVpcLimitNamesToType(request *DescribeVpcLimitNamesToTypeRequest) (response *DescribeVpcLimitNamesToTypeResponse, err error) {
	if request == nil {
		request = NewDescribeVpcLimitNamesToTypeRequest()
	}
	response = NewDescribeVpcLimitNamesToTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcLimitsRequest() (request *DescribeVpcLimitsRequest) {
	request = &DescribeVpcLimitsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcLimits")
	return
}

func NewDescribeVpcLimitsResponse() (response *DescribeVpcLimitsResponse) {
	response = &DescribeVpcLimitsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vpc配额信息列表
func (c *Client) DescribeVpcLimits(request *DescribeVpcLimitsRequest) (response *DescribeVpcLimitsResponse, err error) {
	if request == nil {
		request = NewDescribeVpcLimitsRequest()
	}
	response = NewDescribeVpcLimitsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUnderlayNatGatewayRuleRequest() (request *CreateUnderlayNatGatewayRuleRequest) {
	request = &CreateUnderlayNatGatewayRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateUnderlayNatGatewayRule")
	return
}

func NewCreateUnderlayNatGatewayRuleResponse() (response *CreateUnderlayNatGatewayRuleResponse) {
	response = &CreateUnderlayNatGatewayRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建underlay natgateway规则
func (c *Client) CreateUnderlayNatGatewayRule(request *CreateUnderlayNatGatewayRuleRequest) (response *CreateUnderlayNatGatewayRuleResponse, err error) {
	if request == nil {
		request = NewCreateUnderlayNatGatewayRuleRequest()
	}
	response = NewCreateUnderlayNatGatewayRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewayGroupRequest() (request *DeleteGatewayGroupRequest) {
	request = &DeleteGatewayGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteGatewayGroup")
	return
}

func NewDeleteGatewayGroupResponse() (response *DeleteGatewayGroupResponse) {
	response = &DeleteGatewayGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关集群组
func (c *Client) DeleteGatewayGroup(request *DeleteGatewayGroupRequest) (response *DeleteGatewayGroupResponse, err error) {
	if request == nil {
		request = NewDeleteGatewayGroupRequest()
	}
	response = NewDeleteGatewayGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClbGatewayTypeNamesRequest() (request *DescribeClbGatewayTypeNamesRequest) {
	request = &DescribeClbGatewayTypeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeClbGatewayTypeNames")
	return
}

func NewDescribeClbGatewayTypeNamesResponse() (response *DescribeClbGatewayTypeNamesResponse) {
	response = &DescribeClbGatewayTypeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取clb网关名称
func (c *Client) DescribeClbGatewayTypeNames(request *DescribeClbGatewayTypeNamesRequest) (response *DescribeClbGatewayTypeNamesResponse, err error) {
	if request == nil {
		request = NewDescribeClbGatewayTypeNamesRequest()
	}
	response = NewDescribeClbGatewayTypeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRtbRoutesRequest() (request *DescribeRtbRoutesRequest) {
	request = &DescribeRtbRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeRtbRoutes")
	return
}

func NewDescribeRtbRoutesResponse() (response *DescribeRtbRoutesResponse) {
	response = &DescribeRtbRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由表路由列表
func (c *Client) DescribeRtbRoutes(request *DescribeRtbRoutesRequest) (response *DescribeRtbRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeRtbRoutesRequest()
	}
	response = NewDescribeRtbRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayNodesRequest() (request *DescribeGatewayNodesRequest) {
	request = &DescribeGatewayNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayNodes")
	return
}

func NewDescribeGatewayNodesResponse() (response *DescribeGatewayNodesResponse) {
	response = &DescribeGatewayNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关节点
func (c *Client) DescribeGatewayNodes(request *DescribeGatewayNodesRequest) (response *DescribeGatewayNodesResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayNodesRequest()
	}
	response = NewDescribeGatewayNodesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGatewayAttributeRequest() (request *ModifyGatewayAttributeRequest) {
	request = &ModifyGatewayAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyGatewayAttribute")
	return
}

func NewModifyGatewayAttributeResponse() (response *ModifyGatewayAttributeResponse) {
	response = &ModifyGatewayAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关集群
func (c *Client) ModifyGatewayAttribute(request *ModifyGatewayAttributeRequest) (response *ModifyGatewayAttributeResponse, err error) {
	if request == nil {
		request = NewModifyGatewayAttributeRequest()
	}
	response = NewModifyGatewayAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNatDetectIpRequest() (request *DeleteNatDetectIpRequest) {
	request = &DeleteNatDetectIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteNatDetectIp")
	return
}

func NewDeleteNatDetectIpResponse() (response *DeleteNatDetectIpResponse) {
	response = &DeleteNatDetectIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除NAT网关外网探测地址池。
func (c *Client) DeleteNatDetectIp(request *DeleteNatDetectIpRequest) (response *DeleteNatDetectIpResponse, err error) {
	if request == nil {
		request = NewDeleteNatDetectIpRequest()
	}
	response = NewDeleteNatDetectIpResponse()
	err = c.Send(request, response)
	return
}

func NewGetAclAndItemRequest() (request *GetAclAndItemRequest) {
	request = &GetAclAndItemRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetAclAndItem")
	return
}

func NewGetAclAndItemResponse() (response *GetAclAndItemResponse) {
	response = &GetAclAndItemResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ACL规则信息
func (c *Client) GetAclAndItem(request *GetAclAndItemRequest) (response *GetAclAndItemResponse, err error) {
	if request == nil {
		request = NewGetAclAndItemRequest()
	}
	response = NewGetAclAndItemResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayInstallConfRequest() (request *DescribeGatewayInstallConfRequest) {
	request = &DescribeGatewayInstallConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayInstallConf")
	return
}

func NewDescribeGatewayInstallConfResponse() (response *DescribeGatewayInstallConfResponse) {
	response = &DescribeGatewayInstallConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取一键开局部署模板
func (c *Client) DescribeGatewayInstallConf(request *DescribeGatewayInstallConfRequest) (response *DescribeGatewayInstallConfResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayInstallConfRequest()
	}
	response = NewDescribeGatewayInstallConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayInstallStateRequest() (request *DescribeGatewayInstallStateRequest) {
	request = &DescribeGatewayInstallStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayInstallState")
	return
}

func NewDescribeGatewayInstallStateResponse() (response *DescribeGatewayInstallStateResponse) {
	response = &DescribeGatewayInstallStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关自动化部署状态
func (c *Client) DescribeGatewayInstallState(request *DescribeGatewayInstallStateRequest) (response *DescribeGatewayInstallStateResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayInstallStateRequest()
	}
	response = NewDescribeGatewayInstallStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayVersionsRequest() (request *DescribeGatewayVersionsRequest) {
	request = &DescribeGatewayVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGatewayVersions")
	return
}

func NewDescribeGatewayVersionsResponse() (response *DescribeGatewayVersionsResponse) {
	response = &DescribeGatewayVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网关安装包版本
func (c *Client) DescribeGatewayVersions(request *DescribeGatewayVersionsRequest) (response *DescribeGatewayVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayVersionsRequest()
	}
	response = NewDescribeGatewayVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGatewayGroupAttributeRequest() (request *ModifyGatewayGroupAttributeRequest) {
	request = &ModifyGatewayGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyGatewayGroupAttribute")
	return
}

func NewModifyGatewayGroupAttributeResponse() (response *ModifyGatewayGroupAttributeResponse) {
	response = &ModifyGatewayGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改创建网关集群组
func (c *Client) ModifyGatewayGroupAttribute(request *ModifyGatewayGroupAttributeRequest) (response *ModifyGatewayGroupAttributeResponse, err error) {
	if request == nil {
		request = NewModifyGatewayGroupAttributeRequest()
	}
	response = NewModifyGatewayGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeNfvGwNodeRequest() (request *UpgradeNfvGwNodeRequest) {
	request = &UpgradeNfvGwNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "UpgradeNfvGwNode")
	return
}

func NewUpgradeNfvGwNodeResponse() (response *UpgradeNfvGwNodeResponse) {
	response = &UpgradeNfvGwNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 节点升级
func (c *Client) UpgradeNfvGwNode(request *UpgradeNfvGwNodeRequest) (response *UpgradeNfvGwNodeResponse, err error) {
	if request == nil {
		request = NewUpgradeNfvGwNodeRequest()
	}
	response = NewUpgradeNfvGwNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMulticastGroupsRequest() (request *DescribeMulticastGroupsRequest) {
	request = &DescribeMulticastGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeMulticastGroups")
	return
}

func NewDescribeMulticastGroupsResponse() (response *DescribeMulticastGroupsResponse) {
	response = &DescribeMulticastGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组播组信息。
func (c *Client) DescribeMulticastGroups(request *DescribeMulticastGroupsRequest) (response *DescribeMulticastGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeMulticastGroupsRequest()
	}
	response = NewDescribeMulticastGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEniListRequest() (request *DescribeEniListRequest) {
	request = &DescribeEniListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeEniList")
	return
}

func NewDescribeEniListResponse() (response *DescribeEniListResponse) {
	response = &DescribeEniListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性网卡列表
func (c *Client) DescribeEniList(request *DescribeEniListRequest) (response *DescribeEniListResponse, err error) {
	if request == nil {
		request = NewDescribeEniListRequest()
	}
	response = NewDescribeEniListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayRuleRequest() (request *DescribeUnderlayNatGatewayRuleRequest) {
	request = &DescribeUnderlayNatGatewayRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayRule")
	return
}

func NewDescribeUnderlayNatGatewayRuleResponse() (response *DescribeUnderlayNatGatewayRuleResponse) {
	response = &DescribeUnderlayNatGatewayRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取underlay natgateway规则
func (c *Client) DescribeUnderlayNatGatewayRule(request *DescribeUnderlayNatGatewayRuleRequest) (response *DescribeUnderlayNatGatewayRuleResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayRuleRequest()
	}
	response = NewDescribeUnderlayNatGatewayRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVmRequest() (request *ModifyVmRequest) {
	request = &ModifyVmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyVm")
	return
}

func NewModifyVmResponse() (response *ModifyVmResponse) {
	response = &ModifyVmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改VM。
func (c *Client) ModifyVm(request *ModifyVmRequest) (response *ModifyVmResponse, err error) {
	if request == nil {
		request = NewModifyVmRequest()
	}
	response = NewModifyVmResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceRequest() (request *CreateServiceRequest) {
	request = &CreateServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateService")
	return
}

func NewCreateServiceResponse() (response *CreateServiceResponse) {
	response = &CreateServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建物理网络服务
func (c *Client) CreateService(request *CreateServiceRequest) (response *CreateServiceResponse, err error) {
	if request == nil {
		request = NewCreateServiceRequest()
	}
	response = NewCreateServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkInterfacesExRequest() (request *DescribeNetworkInterfacesExRequest) {
	request = &DescribeNetworkInterfacesExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeNetworkInterfacesEx")
	return
}

func NewDescribeNetworkInterfacesExResponse() (response *DescribeNetworkInterfacesExResponse) {
	response = &DescribeNetworkInterfacesExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询弹性网卡列表
func (c *Client) DescribeNetworkInterfacesEx(request *DescribeNetworkInterfacesExRequest) (response *DescribeNetworkInterfacesExResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkInterfacesExRequest()
	}
	response = NewDescribeNetworkInterfacesExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayRequest() (request *CreateGatewayRequest) {
	request = &CreateGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateGateway")
	return
}

func NewCreateGatewayResponse() (response *CreateGatewayResponse) {
	response = &CreateGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网关集群
func (c *Client) CreateGateway(request *CreateGatewayRequest) (response *CreateGatewayResponse, err error) {
	if request == nil {
		request = NewCreateGatewayRequest()
	}
	response = NewCreateGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCvmForSecurityGroupRequest() (request *DescribeCvmForSecurityGroupRequest) {
	request = &DescribeCvmForSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeCvmForSecurityGroup")
	return
}

func NewDescribeCvmForSecurityGroupResponse() (response *DescribeCvmForSecurityGroupResponse) {
	response = &DescribeCvmForSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全组关联实例列表
func (c *Client) DescribeCvmForSecurityGroup(request *DescribeCvmForSecurityGroupRequest) (response *DescribeCvmForSecurityGroupResponse, err error) {
	if request == nil {
		request = NewDescribeCvmForSecurityGroupRequest()
	}
	response = NewDescribeCvmForSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewAdjustNfvGwRequest() (request *AdjustNfvGwRequest) {
	request = &AdjustNfvGwRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "AdjustNfvGw")
	return
}

func NewAdjustNfvGwResponse() (response *AdjustNfvGwResponse) {
	response = &AdjustNfvGwResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调整NfvGw
func (c *Client) AdjustNfvGw(request *AdjustNfvGwRequest) (response *AdjustNfvGwResponse, err error) {
	if request == nil {
		request = NewAdjustNfvGwRequest()
	}
	response = NewAdjustNfvGwResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDetectInstallationRequest() (request *CreateDetectInstallationRequest) {
	request = &CreateDetectInstallationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateDetectInstallation")
	return
}

func NewCreateDetectInstallationResponse() (response *CreateDetectInstallationResponse) {
	response = &CreateDetectInstallationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateDetectInstallation
func (c *Client) CreateDetectInstallation(request *CreateDetectInstallationRequest) (response *CreateDetectInstallationResponse, err error) {
	if request == nil {
		request = NewCreateDetectInstallationRequest()
	}
	response = NewCreateDetectInstallationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskResultRequest() (request *DescribeTaskResultRequest) {
	request = &DescribeTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeTaskResult")
	return
}

func NewDescribeTaskResultResponse() (response *DescribeTaskResultResponse) {
	response = &DescribeTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务进度
func (c *Client) DescribeTaskResult(request *DescribeTaskResultRequest) (response *DescribeTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeTaskResultRequest()
	}
	response = NewDescribeTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUnderlayNatGatewayWanInfoRequest() (request *DeleteUnderlayNatGatewayWanInfoRequest) {
	request = &DeleteUnderlayNatGatewayWanInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteUnderlayNatGatewayWanInfo")
	return
}

func NewDeleteUnderlayNatGatewayWanInfoResponse() (response *DeleteUnderlayNatGatewayWanInfoResponse) {
	response = &DeleteUnderlayNatGatewayWanInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除underlay natgateway运营商信息
func (c *Client) DeleteUnderlayNatGatewayWanInfo(request *DeleteUnderlayNatGatewayWanInfoRequest) (response *DeleteUnderlayNatGatewayWanInfoResponse, err error) {
	if request == nil {
		request = NewDeleteUnderlayNatGatewayWanInfoRequest()
	}
	response = NewDeleteUnderlayNatGatewayWanInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayInstallationRequest() (request *CreateGatewayInstallationRequest) {
	request = &CreateGatewayInstallationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateGatewayInstallation")
	return
}

func NewCreateGatewayInstallationResponse() (response *CreateGatewayInstallationResponse) {
	response = &CreateGatewayInstallationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网关部署
func (c *Client) CreateGatewayInstallation(request *CreateGatewayInstallationRequest) (response *CreateGatewayInstallationResponse, err error) {
	if request == nil {
		request = NewCreateGatewayInstallationRequest()
	}
	response = NewCreateGatewayInstallationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayCarriorInfoRequest() (request *DescribeUnderlayNatGatewayCarriorInfoRequest) {
	request = &DescribeUnderlayNatGatewayCarriorInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayCarriorInfo")
	return
}

func NewDescribeUnderlayNatGatewayCarriorInfoResponse() (response *DescribeUnderlayNatGatewayCarriorInfoResponse) {
	response = &DescribeUnderlayNatGatewayCarriorInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeUnderlayNatGatewayCarriorInfo
func (c *Client) DescribeUnderlayNatGatewayCarriorInfo(request *DescribeUnderlayNatGatewayCarriorInfoRequest) (response *DescribeUnderlayNatGatewayCarriorInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayCarriorInfoRequest()
	}
	response = NewDescribeUnderlayNatGatewayCarriorInfoResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductStateInfoRequest() (request *QueryProductStateInfoRequest) {
	request = &QueryProductStateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "QueryProductStateInfo")
	return
}

func NewQueryProductStateInfoResponse() (response *QueryProductStateInfoResponse) {
	response = &QueryProductStateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// VPC产品状态信息查询接口
func (c *Client) QueryProductStateInfo(request *QueryProductStateInfoRequest) (response *QueryProductStateInfoResponse, err error) {
	if request == nil {
		request = NewQueryProductStateInfoRequest()
	}
	response = NewQueryProductStateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGatewaysRequest() (request *DescribeUserGatewaysRequest) {
	request = &DescribeUserGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUserGateways")
	return
}

func NewDescribeUserGatewaysResponse() (response *DescribeUserGatewaysResponse) {
	response = &DescribeUserGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vpn对端网关列表
func (c *Client) DescribeUserGateways(request *DescribeUserGatewaysRequest) (response *DescribeUserGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeUserGatewaysRequest()
	}
	response = NewDescribeUserGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewGetNfvGwNodeDetailRequest() (request *GetNfvGwNodeDetailRequest) {
	request = &GetNfvGwNodeDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvGwNodeDetail")
	return
}

func NewGetNfvGwNodeDetailResponse() (response *GetNfvGwNodeDetailResponse) {
	response = &GetNfvGwNodeDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Nfv网关节点的详细信息
func (c *Client) GetNfvGwNodeDetail(request *GetNfvGwNodeDetailRequest) (response *GetNfvGwNodeDetailResponse, err error) {
	if request == nil {
		request = NewGetNfvGwNodeDetailRequest()
	}
	response = NewGetNfvGwNodeDetailResponse()
	err = c.Send(request, response)
	return
}

func NewChangeNfvGwNodeRequest() (request *ChangeNfvGwNodeRequest) {
	request = &ChangeNfvGwNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ChangeNfvGwNode")
	return
}

func NewChangeNfvGwNodeResponse() (response *ChangeNfvGwNodeResponse) {
	response = &ChangeNfvGwNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 改变Nfv网关节点
func (c *Client) ChangeNfvGwNode(request *ChangeNfvGwNodeRequest) (response *ChangeNfvGwNodeResponse, err error) {
	if request == nil {
		request = NewChangeNfvGwNodeRequest()
	}
	response = NewChangeNfvGwNodeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGatewayWeightRequest() (request *ModifyGatewayWeightRequest) {
	request = &ModifyGatewayWeightRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyGatewayWeight")
	return
}

func NewModifyGatewayWeightResponse() (response *ModifyGatewayWeightResponse) {
	response = &ModifyGatewayWeightResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关集群权重，目前只支持vpcgw
func (c *Client) ModifyGatewayWeight(request *ModifyGatewayWeightRequest) (response *ModifyGatewayWeightResponse, err error) {
	if request == nil {
		request = NewModifyGatewayWeightRequest()
	}
	response = NewModifyGatewayWeightResponse()
	err = c.Send(request, response)
	return
}

func NewAdjustNfvGwNodeRequest() (request *AdjustNfvGwNodeRequest) {
	request = &AdjustNfvGwNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "AdjustNfvGwNode")
	return
}

func NewAdjustNfvGwNodeResponse() (response *AdjustNfvGwNodeResponse) {
	response = &AdjustNfvGwNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调整Nfv网关节点
func (c *Client) AdjustNfvGwNode(request *AdjustNfvGwNodeRequest) (response *AdjustNfvGwNodeResponse, err error) {
	if request == nil {
		request = NewAdjustNfvGwNodeRequest()
	}
	response = NewAdjustNfvGwNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClbGatewayInstallRecordRequest() (request *DeleteClbGatewayInstallRecordRequest) {
	request = &DeleteClbGatewayInstallRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteClbGatewayInstallRecord")
	return
}

func NewDeleteClbGatewayInstallRecordResponse() (response *DeleteClbGatewayInstallRecordResponse) {
	response = &DeleteClbGatewayInstallRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除clb网关安装记录
func (c *Client) DeleteClbGatewayInstallRecord(request *DeleteClbGatewayInstallRecordRequest) (response *DeleteClbGatewayInstallRecordResponse, err error) {
	if request == nil {
		request = NewDeleteClbGatewayInstallRecordRequest()
	}
	response = NewDeleteClbGatewayInstallRecordResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNatDetectIpRequest() (request *CreateNatDetectIpRequest) {
	request = &CreateNatDetectIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateNatDetectIp")
	return
}

func NewCreateNatDetectIpResponse() (response *CreateNatDetectIpResponse) {
	response = &CreateNatDetectIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建natgw探测IP。
func (c *Client) CreateNatDetectIp(request *CreateNatDetectIpRequest) (response *CreateNatDetectIpResponse, err error) {
	if request == nil {
		request = NewCreateNatDetectIpRequest()
	}
	response = NewCreateNatDetectIpResponse()
	err = c.Send(request, response)
	return
}

func NewCreateJnsGatewayServiceRequest() (request *CreateJnsGatewayServiceRequest) {
	request = &CreateJnsGatewayServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateJnsGatewayService")
	return
}

func NewCreateJnsGatewayServiceResponse() (response *CreateJnsGatewayServiceResponse) {
	response = &CreateJnsGatewayServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建私有网络映射服务
func (c *Client) CreateJnsGatewayService(request *CreateJnsGatewayServiceRequest) (response *CreateJnsGatewayServiceResponse, err error) {
	if request == nil {
		request = NewCreateJnsGatewayServiceRequest()
	}
	response = NewCreateJnsGatewayServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcPeerCrossRegionRequest() (request *DescribeVpcPeerCrossRegionRequest) {
	request = &DescribeVpcPeerCrossRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcPeerCrossRegion")
	return
}

func NewDescribeVpcPeerCrossRegionResponse() (response *DescribeVpcPeerCrossRegionResponse) {
	response = &DescribeVpcPeerCrossRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询对等连接有效的地域线路
func (c *Client) DescribeVpcPeerCrossRegion(request *DescribeVpcPeerCrossRegionRequest) (response *DescribeVpcPeerCrossRegionResponse, err error) {
	if request == nil {
		request = NewDescribeVpcPeerCrossRegionRequest()
	}
	response = NewDescribeVpcPeerCrossRegionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIpv6PrefixRequest() (request *ModifyIpv6PrefixRequest) {
	request = &ModifyIpv6PrefixRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyIpv6Prefix")
	return
}

func NewModifyIpv6PrefixResponse() (response *ModifyIpv6PrefixResponse) {
	response = &ModifyIpv6PrefixResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置网段裂解掩码范围。
func (c *Client) ModifyIpv6Prefix(request *ModifyIpv6PrefixRequest) (response *ModifyIpv6PrefixResponse, err error) {
	if request == nil {
		request = NewModifyIpv6PrefixRequest()
	}
	response = NewModifyIpv6PrefixResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDirectConnectGatewaysExRequest() (request *DescribeDirectConnectGatewaysExRequest) {
	request = &DescribeDirectConnectGatewaysExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeDirectConnectGatewaysEx")
	return
}

func NewDescribeDirectConnectGatewaysExResponse() (response *DescribeDirectConnectGatewaysExResponse) {
	response = &DescribeDirectConnectGatewaysExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线网关列表
func (c *Client) DescribeDirectConnectGatewaysEx(request *DescribeDirectConnectGatewaysExRequest) (response *DescribeDirectConnectGatewaysExResponse, err error) {
	if request == nil {
		request = NewDescribeDirectConnectGatewaysExRequest()
	}
	response = NewDescribeDirectConnectGatewaysExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRtbSubnetsRequest() (request *DescribeRtbSubnetsRequest) {
	request = &DescribeRtbSubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeRtbSubnets")
	return
}

func NewDescribeRtbSubnetsResponse() (response *DescribeRtbSubnetsResponse) {
	response = &DescribeRtbSubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查路由表对应的全量子网
func (c *Client) DescribeRtbSubnets(request *DescribeRtbSubnetsRequest) (response *DescribeRtbSubnetsResponse, err error) {
	if request == nil {
		request = NewDescribeRtbSubnetsRequest()
	}
	response = NewDescribeRtbSubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayRuleWanIpsRequest() (request *DescribeUnderlayNatGatewayRuleWanIpsRequest) {
	request = &DescribeUnderlayNatGatewayRuleWanIpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayRuleWanIps")
	return
}

func NewDescribeUnderlayNatGatewayRuleWanIpsResponse() (response *DescribeUnderlayNatGatewayRuleWanIpsResponse) {
	response = &DescribeUnderlayNatGatewayRuleWanIpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取underlay natgateway规则对应运营端外网ip
func (c *Client) DescribeUnderlayNatGatewayRuleWanIps(request *DescribeUnderlayNatGatewayRuleWanIpsRequest) (response *DescribeUnderlayNatGatewayRuleWanIpsResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayRuleWanIpsRequest()
	}
	response = NewDescribeUnderlayNatGatewayRuleWanIpsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnderlayNatGatewayWanInfoAndWanIpRequest() (request *DescribeUnderlayNatGatewayWanInfoAndWanIpRequest) {
	request = &DescribeUnderlayNatGatewayWanInfoAndWanIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeUnderlayNatGatewayWanInfoAndWanIp")
	return
}

func NewDescribeUnderlayNatGatewayWanInfoAndWanIpResponse() (response *DescribeUnderlayNatGatewayWanInfoAndWanIpResponse) {
	response = &DescribeUnderlayNatGatewayWanInfoAndWanIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取underlay natgateway运营商信息和对应的外网ip
func (c *Client) DescribeUnderlayNatGatewayWanInfoAndWanIp(request *DescribeUnderlayNatGatewayWanInfoAndWanIpRequest) (response *DescribeUnderlayNatGatewayWanInfoAndWanIpResponse, err error) {
	if request == nil {
		request = NewDescribeUnderlayNatGatewayWanInfoAndWanIpRequest()
	}
	response = NewDescribeUnderlayNatGatewayWanInfoAndWanIpResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProtectPolicyIpRequest() (request *ModifyProtectPolicyIpRequest) {
	request = &ModifyProtectPolicyIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyProtectPolicyIp")
	return
}

func NewModifyProtectPolicyIpResponse() (response *ModifyProtectPolicyIpResponse) {
	response = &ModifyProtectPolicyIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关过载保护信息。
func (c *Client) ModifyProtectPolicyIp(request *ModifyProtectPolicyIpRequest) (response *ModifyProtectPolicyIpResponse, err error) {
	if request == nil {
		request = NewModifyProtectPolicyIpRequest()
	}
	response = NewModifyProtectPolicyIpResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJnsGatewayServicesRequest() (request *DescribeJnsGatewayServicesRequest) {
	request = &DescribeJnsGatewayServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeJnsGatewayServices")
	return
}

func NewDescribeJnsGatewayServicesResponse() (response *DescribeJnsGatewayServicesResponse) {
	response = &DescribeJnsGatewayServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询私有网络映射服务
func (c *Client) DescribeJnsGatewayServices(request *DescribeJnsGatewayServicesRequest) (response *DescribeJnsGatewayServicesResponse, err error) {
	if request == nil {
		request = NewDescribeJnsGatewayServicesRequest()
	}
	response = NewDescribeJnsGatewayServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNatDetectIpsRequest() (request *DescribeNatDetectIpsRequest) {
	request = &DescribeNatDetectIpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeNatDetectIps")
	return
}

func NewDescribeNatDetectIpsResponse() (response *DescribeNatDetectIpsResponse) {
	response = &DescribeNatDetectIpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取NAT网关外网探测地址池。
func (c *Client) DescribeNatDetectIps(request *DescribeNatDetectIpsRequest) (response *DescribeNatDetectIpsResponse, err error) {
	if request == nil {
		request = NewDescribeNatDetectIpsRequest()
	}
	response = NewDescribeNatDetectIpsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewayInstallRecordsRequest() (request *DeleteGatewayInstallRecordsRequest) {
	request = &DeleteGatewayInstallRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteGatewayInstallRecords")
	return
}

func NewDeleteGatewayInstallRecordsResponse() (response *DeleteGatewayInstallRecordsResponse) {
	response = &DeleteGatewayInstallRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关部署记录
func (c *Client) DeleteGatewayInstallRecords(request *DeleteGatewayInstallRecordsRequest) (response *DeleteGatewayInstallRecordsResponse, err error) {
	if request == nil {
		request = NewDeleteGatewayInstallRecordsRequest()
	}
	response = NewDeleteGatewayInstallRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIpV6SubnetsRequest() (request *DeleteIpV6SubnetsRequest) {
	request = &DeleteIpV6SubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteIpV6Subnets")
	return
}

func NewDeleteIpV6SubnetsResponse() (response *DeleteIpV6SubnetsResponse) {
	response = &DeleteIpV6SubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除VPC IPv6地址段。
func (c *Client) DeleteIpV6Subnets(request *DeleteIpV6SubnetsRequest) (response *DeleteIpV6SubnetsResponse, err error) {
	if request == nil {
		request = NewDeleteIpV6SubnetsRequest()
	}
	response = NewDeleteIpV6SubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpV6SubnetsRequest() (request *DescribeIpV6SubnetsRequest) {
	request = &DescribeIpV6SubnetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeIpV6Subnets")
	return
}

func NewDescribeIpV6SubnetsResponse() (response *DescribeIpV6SubnetsResponse) {
	response = &DescribeIpV6SubnetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取VPC IPv6地址段。
func (c *Client) DescribeIpV6Subnets(request *DescribeIpV6SubnetsRequest) (response *DescribeIpV6SubnetsResponse, err error) {
	if request == nil {
		request = NewDescribeIpV6SubnetsRequest()
	}
	response = NewDescribeIpV6SubnetsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMulticastGroupRequest() (request *ModifyMulticastGroupRequest) {
	request = &ModifyMulticastGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyMulticastGroup")
	return
}

func NewModifyMulticastGroupResponse() (response *ModifyMulticastGroupResponse) {
	response = &ModifyMulticastGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改组播组属性。
func (c *Client) ModifyMulticastGroup(request *ModifyMulticastGroupRequest) (response *ModifyMulticastGroupResponse, err error) {
	if request == nil {
		request = NewModifyMulticastGroupRequest()
	}
	response = NewModifyMulticastGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewaysRequest() (request *DescribeGatewaysRequest) {
	request = &DescribeGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeGateways")
	return
}

func NewDescribeGatewaysResponse() (response *DescribeGatewaysResponse) {
	response = &DescribeGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关
func (c *Client) DescribeGateways(request *DescribeGatewaysRequest) (response *DescribeGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeGatewaysRequest()
	}
	response = NewDescribeGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewProceedNfvGwLogStepRequest() (request *ProceedNfvGwLogStepRequest) {
	request = &ProceedNfvGwLogStepRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ProceedNfvGwLogStep")
	return
}

func NewProceedNfvGwLogStepResponse() (response *ProceedNfvGwLogStepResponse) {
	response = &ProceedNfvGwLogStepResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 继续任务
func (c *Client) ProceedNfvGwLogStep(request *ProceedNfvGwLogStepRequest) (response *ProceedNfvGwLogStepResponse, err error) {
	if request == nil {
		request = NewProceedNfvGwLogStepRequest()
	}
	response = NewProceedNfvGwLogStepResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubnetExRequest() (request *DescribeSubnetExRequest) {
	request = &DescribeSubnetExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSubnetEx")
	return
}

func NewDescribeSubnetExResponse() (response *DescribeSubnetExResponse) {
	response = &DescribeSubnetExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子网列表
func (c *Client) DescribeSubnetEx(request *DescribeSubnetExRequest) (response *DescribeSubnetExResponse, err error) {
	if request == nil {
		request = NewDescribeSubnetExRequest()
	}
	response = NewDescribeSubnetExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDetectHostStateRequest() (request *DescribeDetectHostStateRequest) {
	request = &DescribeDetectHostStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeDetectHostState")
	return
}

func NewDescribeDetectHostStateResponse() (response *DescribeDetectHostStateResponse) {
	response = &DescribeDetectHostStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取探测机状态
func (c *Client) DescribeDetectHostState(request *DescribeDetectHostStateRequest) (response *DescribeDetectHostStateResponse, err error) {
	if request == nil {
		request = NewDescribeDetectHostStateRequest()
	}
	response = NewDescribeDetectHostStateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServicesRequest() (request *DescribeServicesRequest) {
	request = &DescribeServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeServices")
	return
}

func NewDescribeServicesResponse() (response *DescribeServicesResponse) {
	response = &DescribeServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物理网络服务映射列表
func (c *Client) DescribeServices(request *DescribeServicesRequest) (response *DescribeServicesResponse, err error) {
	if request == nil {
		request = NewDescribeServicesRequest()
	}
	response = NewDescribeServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupPolicyRequest() (request *DescribeSecurityGroupPolicyRequest) {
	request = &DescribeSecurityGroupPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSecurityGroupPolicy")
	return
}

func NewDescribeSecurityGroupPolicyResponse() (response *DescribeSecurityGroupPolicyResponse) {
	response = &DescribeSecurityGroupPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全组策略
func (c *Client) DescribeSecurityGroupPolicy(request *DescribeSecurityGroupPolicyRequest) (response *DescribeSecurityGroupPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupPolicyRequest()
	}
	response = NewDescribeSecurityGroupPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointServiceWhiteListRequest() (request *DescribeVpcEndPointServiceWhiteListRequest) {
	request = &DescribeVpcEndPointServiceWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcEndPointServiceWhiteList")
	return
}

func NewDescribeVpcEndPointServiceWhiteListResponse() (response *DescribeVpcEndPointServiceWhiteListResponse) {
	response = &DescribeVpcEndPointServiceWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询终端节点服务的服务白名单。
func (c *Client) DescribeVpcEndPointServiceWhiteList(request *DescribeVpcEndPointServiceWhiteListRequest) (response *DescribeVpcEndPointServiceWhiteListResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointServiceWhiteListRequest()
	}
	response = NewDescribeVpcEndPointServiceWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDetectInstallRecordsRequest() (request *DeleteDetectInstallRecordsRequest) {
	request = &DeleteDetectInstallRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteDetectInstallRecords")
	return
}

func NewDeleteDetectInstallRecordsResponse() (response *DeleteDetectInstallRecordsResponse) {
	response = &DeleteDetectInstallRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除探测机部署记录
func (c *Client) DeleteDetectInstallRecords(request *DeleteDetectInstallRecordsRequest) (response *DeleteDetectInstallRecordsResponse, err error) {
	if request == nil {
		request = NewDeleteDetectInstallRecordsRequest()
	}
	response = NewDeleteDetectInstallRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewGetNfvGwTypesRequest() (request *GetNfvGwTypesRequest) {
	request = &GetNfvGwTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvGwTypes")
	return
}

func NewGetNfvGwTypesResponse() (response *GetNfvGwTypesResponse) {
	response = &GetNfvGwTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Nfv网关的类型
func (c *Client) GetNfvGwTypes(request *GetNfvGwTypesRequest) (response *GetNfvGwTypesResponse, err error) {
	if request == nil {
		request = NewGetNfvGwTypesRequest()
	}
	response = NewGetNfvGwTypesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDetectHostWeightRequest() (request *ModifyDetectHostWeightRequest) {
	request = &ModifyDetectHostWeightRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyDetectHostWeight")
	return
}

func NewModifyDetectHostWeightResponse() (response *ModifyDetectHostWeightResponse) {
	response = &ModifyDetectHostWeightResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新探测机权重
func (c *Client) ModifyDetectHostWeight(request *ModifyDetectHostWeightRequest) (response *ModifyDetectHostWeightResponse, err error) {
	if request == nil {
		request = NewModifyDetectHostWeightRequest()
	}
	response = NewModifyDetectHostWeightResponse()
	err = c.Send(request, response)
	return
}

func NewGetNfvSwitchGroupRequest() (request *GetNfvSwitchGroupRequest) {
	request = &GetNfvSwitchGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvSwitchGroup")
	return
}

func NewGetNfvSwitchGroupResponse() (response *GetNfvSwitchGroupResponse) {
	response = &GetNfvSwitchGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取nfv交换机组信息
func (c *Client) GetNfvSwitchGroup(request *GetNfvSwitchGroupRequest) (response *GetNfvSwitchGroupResponse, err error) {
	if request == nil {
		request = NewGetNfvSwitchGroupRequest()
	}
	response = NewGetNfvSwitchGroupResponse()
	err = c.Send(request, response)
	return
}

func NewAddUnderlayNatGatewayWanIpRequest() (request *AddUnderlayNatGatewayWanIpRequest) {
	request = &AddUnderlayNatGatewayWanIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "AddUnderlayNatGatewayWanIp")
	return
}

func NewAddUnderlayNatGatewayWanIpResponse() (response *AddUnderlayNatGatewayWanIpResponse) {
	response = &AddUnderlayNatGatewayWanIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增underlay natgateway运营商外网ip
func (c *Client) AddUnderlayNatGatewayWanIp(request *AddUnderlayNatGatewayWanIpRequest) (response *AddUnderlayNatGatewayWanIpResponse, err error) {
	if request == nil {
		request = NewAddUnderlayNatGatewayWanIpRequest()
	}
	response = NewAddUnderlayNatGatewayWanIpResponse()
	err = c.Send(request, response)
	return
}

func NewGetNfvCvmInstanceTypeRequest() (request *GetNfvCvmInstanceTypeRequest) {
	request = &GetNfvCvmInstanceTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "GetNfvCvmInstanceType")
	return
}

func NewGetNfvCvmInstanceTypeResponse() (response *GetNfvCvmInstanceTypeResponse) {
	response = &GetNfvCvmInstanceTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取支持的NFV母机类型
func (c *Client) GetNfvCvmInstanceType(request *GetNfvCvmInstanceTypeRequest) (response *GetNfvCvmInstanceTypeResponse, err error) {
	if request == nil {
		request = NewGetNfvCvmInstanceTypeRequest()
	}
	response = NewGetNfvCvmInstanceTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest() (request *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest) {
	request = &ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ModifyUnderlayNatGatewayWanInfoAndWanIpAttribute")
	return
}

func NewModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse() (response *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse) {
	response = &ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改underlay natgateway运营商信息和对应的外网ip
func (c *Client) ModifyUnderlayNatGatewayWanInfoAndWanIpAttribute(request *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest) (response *ModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse, err error) {
	if request == nil {
		request = NewModifyUnderlayNatGatewayWanInfoAndWanIpAttributeRequest()
	}
	response = NewModifyUnderlayNatGatewayWanInfoAndWanIpAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewAddUnderlayNatGatewayRuleRipRequest() (request *AddUnderlayNatGatewayRuleRipRequest) {
	request = &AddUnderlayNatGatewayRuleRipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "AddUnderlayNatGatewayRuleRip")
	return
}

func NewAddUnderlayNatGatewayRuleRipResponse() (response *AddUnderlayNatGatewayRuleRipResponse) {
	response = &AddUnderlayNatGatewayRuleRipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增underlay natgateway规则对应rip
func (c *Client) AddUnderlayNatGatewayRuleRip(request *AddUnderlayNatGatewayRuleRipRequest) (response *AddUnderlayNatGatewayRuleRipResponse, err error) {
	if request == nil {
		request = NewAddUnderlayNatGatewayRuleRipRequest()
	}
	response = NewAddUnderlayNatGatewayRuleRipResponse()
	err = c.Send(request, response)
	return
}

func NewRebootNfvGwLogStepRequest() (request *RebootNfvGwLogStepRequest) {
	request = &RebootNfvGwLogStepRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "RebootNfvGwLogStep")
	return
}

func NewRebootNfvGwLogStepResponse() (response *RebootNfvGwLogStepResponse) {
	response = &RebootNfvGwLogStepResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新启动Nfv网关日志步骤
func (c *Client) RebootNfvGwLogStep(request *RebootNfvGwLogStepRequest) (response *RebootNfvGwLogStepResponse, err error) {
	if request == nil {
		request = NewRebootNfvGwLogStepRequest()
	}
	response = NewRebootNfvGwLogStepResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNfvGwNameRequest() (request *UpdateNfvGwNameRequest) {
	request = &UpdateNfvGwNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "UpdateNfvGwName")
	return
}

func NewUpdateNfvGwNameResponse() (response *UpdateNfvGwNameResponse) {
	response = &UpdateNfvGwNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Nfv网关名
func (c *Client) UpdateNfvGwName(request *UpdateNfvGwNameRequest) (response *UpdateNfvGwNameResponse, err error) {
	if request == nil {
		request = NewUpdateNfvGwNameRequest()
	}
	response = NewUpdateNfvGwNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecurityGroupRequest() (request *DescribeSecurityGroupRequest) {
	request = &DescribeSecurityGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeSecurityGroup")
	return
}

func NewDescribeSecurityGroupResponse() (response *DescribeSecurityGroupResponse) {
	response = &DescribeSecurityGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取安全组列表
func (c *Client) DescribeSecurityGroup(request *DescribeSecurityGroupRequest) (response *DescribeSecurityGroupResponse, err error) {
	if request == nil {
		request = NewDescribeSecurityGroupRequest()
	}
	response = NewDescribeSecurityGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcEndPointRequest() (request *DescribeVpcEndPointRequest) {
	request = &DescribeVpcEndPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcEndPoint")
	return
}

func NewDescribeVpcEndPointResponse() (response *DescribeVpcEndPointResponse) {
	response = &DescribeVpcEndPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询终端节点详情。
func (c *Client) DescribeVpcEndPoint(request *DescribeVpcEndPointRequest) (response *DescribeVpcEndPointResponse, err error) {
	if request == nil {
		request = NewDescribeVpcEndPointRequest()
	}
	response = NewDescribeVpcEndPointResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVpcGlobalExtendCidrRequest() (request *DescribeVpcGlobalExtendCidrRequest) {
	request = &DescribeVpcGlobalExtendCidrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVpcGlobalExtendCidr")
	return
}

func NewDescribeVpcGlobalExtendCidrResponse() (response *DescribeVpcGlobalExtendCidrResponse) {
	response = &DescribeVpcGlobalExtendCidrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeVpcGlobalExtendCidr
func (c *Client) DescribeVpcGlobalExtendCidr(request *DescribeVpcGlobalExtendCidrRequest) (response *DescribeVpcGlobalExtendCidrResponse, err error) {
	if request == nil {
		request = NewDescribeVpcGlobalExtendCidrRequest()
	}
	response = NewDescribeVpcGlobalExtendCidrResponse()
	err = c.Send(request, response)
	return
}

func NewArrayNfvGwRequest() (request *ArrayNfvGwRequest) {
	request = &ArrayNfvGwRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "ArrayNfvGw")
	return
}

func NewArrayNfvGwResponse() (response *ArrayNfvGwResponse) {
	response = &ArrayNfvGwResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ArrayNfvGw
func (c *Client) ArrayNfvGw(request *ArrayNfvGwRequest) (response *ArrayNfvGwResponse, err error) {
	if request == nil {
		request = NewArrayNfvGwRequest()
	}
	response = NewArrayNfvGwResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayGroupRequest() (request *CreateGatewayGroupRequest) {
	request = &CreateGatewayGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateGatewayGroup")
	return
}

func NewCreateGatewayGroupResponse() (response *CreateGatewayGroupResponse) {
	response = &CreateGatewayGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网关集群组，目前只支持vpcgw，Type填1
func (c *Client) CreateGatewayGroup(request *CreateGatewayGroupRequest) (response *CreateGatewayGroupResponse, err error) {
	if request == nil {
		request = NewCreateGatewayGroupRequest()
	}
	response = NewCreateGatewayGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBasicNetsRequest() (request *CreateBasicNetsRequest) {
	request = &CreateBasicNetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateBasicNets")
	return
}

func NewCreateBasicNetsResponse() (response *CreateBasicNetsResponse) {
	response = &CreateBasicNetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创建vpc0网段
func (c *Client) CreateBasicNets(request *CreateBasicNetsRequest) (response *CreateBasicNetsResponse, err error) {
	if request == nil {
		request = NewCreateBasicNetsRequest()
	}
	response = NewCreateBasicNetsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceRoutesRequest() (request *DeleteServiceRoutesRequest) {
	request = &DeleteServiceRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DeleteServiceRoutes")
	return
}

func NewDeleteServiceRoutesResponse() (response *DeleteServiceRoutesResponse) {
	response = &DeleteServiceRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除物理网络服务路由
func (c *Client) DeleteServiceRoutes(request *DeleteServiceRoutesRequest) (response *DeleteServiceRoutesResponse, err error) {
	if request == nil {
		request = NewDeleteServiceRoutesRequest()
	}
	response = NewDeleteServiceRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewSwitchGatewayRequest() (request *SwitchGatewayRequest) {
	request = &SwitchGatewayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "SwitchGateway")
	return
}

func NewSwitchGatewayResponse() (response *SwitchGatewayResponse) {
	response = &SwitchGatewayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SwitchGateway
func (c *Client) SwitchGateway(request *SwitchGatewayRequest) (response *SwitchGatewayResponse, err error) {
	if request == nil {
		request = NewSwitchGatewayRequest()
	}
	response = NewSwitchGatewayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIpv6PrefixRequest() (request *DescribeIpv6PrefixRequest) {
	request = &DescribeIpv6PrefixRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeIpv6Prefix")
	return
}

func NewDescribeIpv6PrefixResponse() (response *DescribeIpv6PrefixResponse) {
	response = &DescribeIpv6PrefixResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网段裂解掩码。
func (c *Client) DescribeIpv6Prefix(request *DescribeIpv6PrefixRequest) (response *DescribeIpv6PrefixResponse, err error) {
	if request == nil {
		request = NewDescribeIpv6PrefixRequest()
	}
	response = NewDescribeIpv6PrefixResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDetectTypeNamesRequest() (request *DescribeDetectTypeNamesRequest) {
	request = &DescribeDetectTypeNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeDetectTypeNames")
	return
}

func NewDescribeDetectTypeNamesResponse() (response *DescribeDetectTypeNamesResponse) {
	response = &DescribeDetectTypeNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取探测机类型名字列表
func (c *Client) DescribeDetectTypeNames(request *DescribeDetectTypeNamesRequest) (response *DescribeDetectTypeNamesResponse, err error) {
	if request == nil {
		request = NewDescribeDetectTypeNamesRequest()
	}
	response = NewDescribeDetectTypeNamesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRtbsExRequest() (request *DescribeRtbsExRequest) {
	request = &DescribeRtbsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeRtbsEx")
	return
}

func NewDescribeRtbsExResponse() (response *DescribeRtbsExResponse) {
	response = &DescribeRtbsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由列表
func (c *Client) DescribeRtbsEx(request *DescribeRtbsExRequest) (response *DescribeRtbsExResponse, err error) {
	if request == nil {
		request = NewDescribeRtbsExRequest()
	}
	response = NewDescribeRtbsExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateVpcPeerCrossRegionRequest() (request *CreateVpcPeerCrossRegionRequest) {
	request = &CreateVpcPeerCrossRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "CreateVpcPeerCrossRegion")
	return
}

func NewCreateVpcPeerCrossRegionResponse() (response *CreateVpcPeerCrossRegionResponse) {
	response = &CreateVpcPeerCrossRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建对等连接有效的地域线路
func (c *Client) CreateVpcPeerCrossRegion(request *CreateVpcPeerCrossRegionRequest) (response *CreateVpcPeerCrossRegionResponse, err error) {
	if request == nil {
		request = NewCreateVpcPeerCrossRegionRequest()
	}
	response = NewCreateVpcPeerCrossRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVmRequest() (request *DescribeVmRequest) {
	request = &DescribeVmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "DescribeVm")
	return
}

func NewDescribeVmResponse() (response *DescribeVmResponse) {
	response = &DescribeVmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VPC子机
func (c *Client) DescribeVm(request *DescribeVmRequest) (response *DescribeVmResponse, err error) {
	if request == nil {
		request = NewDescribeVmRequest()
	}
	response = NewDescribeVmResponse()
	err = c.Send(request, response)
	return
}

func NewAddServiceRoutesRequest() (request *AddServiceRoutesRequest) {
	request = &AddServiceRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("opvpc", APIVersion, "AddServiceRoutes")
	return
}

func NewAddServiceRoutesResponse() (response *AddServiceRoutesResponse) {
	response = &AddServiceRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加物理网络服务路由
func (c *Client) AddServiceRoutes(request *AddServiceRoutesRequest) (response *AddServiceRoutesResponse, err error) {
	if request == nil {
		request = NewAddServiceRoutesRequest()
	}
	response = NewAddServiceRoutesResponse()
	err = c.Send(request, response)
	return
}
