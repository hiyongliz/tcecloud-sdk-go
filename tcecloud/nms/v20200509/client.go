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

package v20200509

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-05-09"

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

func NewDescribeDeviceRoleProportionRequest() (request *DescribeDeviceRoleProportionRequest) {
	request = &DescribeDeviceRoleProportionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeDeviceRoleProportion")
	return
}

func NewDescribeDeviceRoleProportionResponse() (response *DescribeDeviceRoleProportionResponse) {
	response = &DescribeDeviceRoleProportionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询设备角色占比信息
func (c *Client) DescribeDeviceRoleProportion(request *DescribeDeviceRoleProportionRequest) (response *DescribeDeviceRoleProportionResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceRoleProportionRequest()
	}
	response = NewDescribeDeviceRoleProportionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkDevicesRequest() (request *ModifyNetworkDevicesRequest) {
	request = &ModifyNetworkDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyNetworkDevices")
	return
}

func NewModifyNetworkDevicesResponse() (response *ModifyNetworkDevicesResponse) {
	response = &ModifyNetworkDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网络设备信息
func (c *Client) ModifyNetworkDevices(request *ModifyNetworkDevicesRequest) (response *ModifyNetworkDevicesResponse, err error) {
	if request == nil {
		request = NewModifyNetworkDevicesRequest()
	}
	response = NewModifyNetworkDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDeviceSyslogKeywordsRequest() (request *DescribeNetDeviceSyslogKeywordsRequest) {
	request = &DescribeNetDeviceSyslogKeywordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDeviceSyslogKeywords")
	return
}

func NewDescribeNetDeviceSyslogKeywordsResponse() (response *DescribeNetDeviceSyslogKeywordsResponse) {
	response = &DescribeNetDeviceSyslogKeywordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网络设备的系统日志关键字列表
func (c *Client) DescribeNetDeviceSyslogKeywords(request *DescribeNetDeviceSyslogKeywordsRequest) (response *DescribeNetDeviceSyslogKeywordsResponse, err error) {
	if request == nil {
		request = NewDescribeNetDeviceSyslogKeywordsRequest()
	}
	response = NewDescribeNetDeviceSyslogKeywordsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpecialIdcLineRequest() (request *ModifySpecialIdcLineRequest) {
	request = &ModifySpecialIdcLineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifySpecialIdcLine")
	return
}

func NewModifySpecialIdcLineResponse() (response *ModifySpecialIdcLineResponse) {
	response = &ModifySpecialIdcLineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改IDC专线
func (c *Client) ModifySpecialIdcLine(request *ModifySpecialIdcLineRequest) (response *ModifySpecialIdcLineResponse, err error) {
	if request == nil {
		request = NewModifySpecialIdcLineRequest()
	}
	response = NewModifySpecialIdcLineResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateIPSegmentsRequest() (request *UpdateIPSegmentsRequest) {
	request = &UpdateIPSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "UpdateIPSegments")
	return
}

func NewUpdateIPSegmentsResponse() (response *UpdateIPSegmentsResponse) {
	response = &UpdateIPSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新IP网段信息
func (c *Client) UpdateIPSegments(request *UpdateIPSegmentsRequest) (response *UpdateIPSegmentsResponse, err error) {
	if request == nil {
		request = NewUpdateIPSegmentsRequest()
	}
	response = NewUpdateIPSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayClustersRequest() (request *CreateGatewayClustersRequest) {
	request = &CreateGatewayClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateGatewayClusters")
	return
}

func NewCreateGatewayClustersResponse() (response *CreateGatewayClustersResponse) {
	response = &CreateGatewayClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加网关集群
func (c *Client) CreateGatewayClusters(request *CreateGatewayClustersRequest) (response *CreateGatewayClustersResponse, err error) {
	if request == nil {
		request = NewCreateGatewayClustersRequest()
	}
	response = NewCreateGatewayClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkRoleDictRequest() (request *DescribeNetworkRoleDictRequest) {
	request = &DescribeNetworkRoleDictRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetworkRoleDict")
	return
}

func NewDescribeNetworkRoleDictResponse() (response *DescribeNetworkRoleDictResponse) {
	response = &DescribeNetworkRoleDictResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备角色中英文字典
func (c *Client) DescribeNetworkRoleDict(request *DescribeNetworkRoleDictRequest) (response *DescribeNetworkRoleDictResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkRoleDictRequest()
	}
	response = NewDescribeNetworkRoleDictResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetDeviceConfigForNmsUpgradeRequest() (request *ModifyNetDeviceConfigForNmsUpgradeRequest) {
	request = &ModifyNetDeviceConfigForNmsUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyNetDeviceConfigForNmsUpgrade")
	return
}

func NewModifyNetDeviceConfigForNmsUpgradeResponse() (response *ModifyNetDeviceConfigForNmsUpgradeResponse) {
	response = &ModifyNetDeviceConfigForNmsUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Nms升级依赖设备配置修改
func (c *Client) ModifyNetDeviceConfigForNmsUpgrade(request *ModifyNetDeviceConfigForNmsUpgradeRequest) (response *ModifyNetDeviceConfigForNmsUpgradeResponse, err error) {
	if request == nil {
		request = NewModifyNetDeviceConfigForNmsUpgradeRequest()
	}
	response = NewModifyNetDeviceConfigForNmsUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIPAddressesRequest() (request *DeleteIPAddressesRequest) {
	request = &DeleteIPAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteIPAddresses")
	return
}

func NewDeleteIPAddressesResponse() (response *DeleteIPAddressesResponse) {
	response = &DeleteIPAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 释放IP地址的占用
func (c *Client) DeleteIPAddresses(request *DeleteIPAddressesRequest) (response *DeleteIPAddressesResponse, err error) {
	if request == nil {
		request = NewDeleteIPAddressesRequest()
	}
	response = NewDeleteIPAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceVendorProportionRequest() (request *DescribeDeviceVendorProportionRequest) {
	request = &DescribeDeviceVendorProportionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeDeviceVendorProportion")
	return
}

func NewDescribeDeviceVendorProportionResponse() (response *DescribeDeviceVendorProportionResponse) {
	response = &DescribeDeviceVendorProportionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询设备厂家占比信息
func (c *Client) DescribeDeviceVendorProportion(request *DescribeDeviceVendorProportionRequest) (response *DescribeDeviceVendorProportionResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceVendorProportionRequest()
	}
	response = NewDescribeDeviceVendorProportionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIdcDirectConnectMetricRequest() (request *DescribeIdcDirectConnectMetricRequest) {
	request = &DescribeIdcDirectConnectMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIdcDirectConnectMetric")
	return
}

func NewDescribeIdcDirectConnectMetricResponse() (response *DescribeIdcDirectConnectMetricResponse) {
	response = &DescribeIdcDirectConnectMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询AZ内IDC物理专线的指标数据
func (c *Client) DescribeIdcDirectConnectMetric(request *DescribeIdcDirectConnectMetricRequest) (response *DescribeIdcDirectConnectMetricResponse, err error) {
	if request == nil {
		request = NewDescribeIdcDirectConnectMetricRequest()
	}
	response = NewDescribeIdcDirectConnectMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerRequest() (request *DescribeServerRequest) {
	request = &DescribeServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeServer")
	return
}

func NewDescribeServerResponse() (response *DescribeServerResponse) {
	response = &DescribeServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务器信息
func (c *Client) DescribeServer(request *DescribeServerRequest) (response *DescribeServerResponse, err error) {
	if request == nil {
		request = NewDescribeServerRequest()
	}
	response = NewDescribeServerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPCIDRSegmentsRequest() (request *DescribeIPCIDRSegmentsRequest) {
	request = &DescribeIPCIDRSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIPCIDRSegments")
	return
}

func NewDescribeIPCIDRSegmentsResponse() (response *DescribeIPCIDRSegmentsResponse) {
	response = &DescribeIPCIDRSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询IPCIDR网段信息
func (c *Client) DescribeIPCIDRSegments(request *DescribeIPCIDRSegmentsRequest) (response *DescribeIPCIDRSegmentsResponse, err error) {
	if request == nil {
		request = NewDescribeIPCIDRSegmentsRequest()
	}
	response = NewDescribeIPCIDRSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkRoleDictRequest() (request *ModifyNetworkRoleDictRequest) {
	request = &ModifyNetworkRoleDictRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyNetworkRoleDict")
	return
}

func NewModifyNetworkRoleDictResponse() (response *ModifyNetworkRoleDictResponse) {
	response = &ModifyNetworkRoleDictResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网络设备角色中英文字典
func (c *Client) ModifyNetworkRoleDict(request *ModifyNetworkRoleDictRequest) (response *ModifyNetworkRoleDictResponse, err error) {
	if request == nil {
		request = NewModifyNetworkRoleDictRequest()
	}
	response = NewModifyNetworkRoleDictResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRacksRequest() (request *DeleteRacksRequest) {
	request = &DeleteRacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteRacks")
	return
}

func NewDeleteRacksResponse() (response *DeleteRacksResponse) {
	response = &DeleteRacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机架
func (c *Client) DeleteRacks(request *DeleteRacksRequest) (response *DeleteRacksResponse, err error) {
	if request == nil {
		request = NewDeleteRacksRequest()
	}
	response = NewDeleteRacksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDeviceRequest() (request *DescribeNetDeviceRequest) {
	request = &DescribeNetDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevice")
	return
}

func NewDescribeNetDeviceResponse() (response *DescribeNetDeviceResponse) {
	response = &DescribeNetDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备信息
func (c *Client) DescribeNetDevice(request *DescribeNetDeviceRequest) (response *DescribeNetDeviceResponse, err error) {
	if request == nil {
		request = NewDescribeNetDeviceRequest()
	}
	response = NewDescribeNetDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkPortsRequest() (request *DeleteNetworkPortsRequest) {
	request = &DeleteNetworkPortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteNetworkPorts")
	return
}

func NewDeleteNetworkPortsResponse() (response *DeleteNetworkPortsResponse) {
	response = &DeleteNetworkPortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络设备端口
func (c *Client) DeleteNetworkPorts(request *DeleteNetworkPortsRequest) (response *DeleteNetworkPortsResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkPortsRequest()
	}
	response = NewDeleteNetworkPortsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
	request = &DescribeTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeTemplate")
	return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
	response = &DescribeTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备配置模板
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateRequest()
	}
	response = NewDescribeTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateTemplate")
	return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加网络设备配置模板
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	if request == nil {
		request = NewCreateTemplateRequest()
	}
	response = NewCreateTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetdeviceRelatedServersRequest() (request *DescribeNetdeviceRelatedServersRequest) {
	request = &DescribeNetdeviceRelatedServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetdeviceRelatedServers")
	return
}

func NewDescribeNetdeviceRelatedServersResponse() (response *DescribeNetdeviceRelatedServersResponse) {
	response = &DescribeNetdeviceRelatedServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备下联服务器
func (c *Client) DescribeNetdeviceRelatedServers(request *DescribeNetdeviceRelatedServersRequest) (response *DescribeNetdeviceRelatedServersResponse, err error) {
	if request == nil {
		request = NewDescribeNetdeviceRelatedServersRequest()
	}
	response = NewDescribeNetdeviceRelatedServersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPAddressesRequest() (request *DescribeIPAddressesRequest) {
	request = &DescribeIPAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIPAddresses")
	return
}

func NewDescribeIPAddressesResponse() (response *DescribeIPAddressesResponse) {
	response = &DescribeIPAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取IP地址
func (c *Client) DescribeIPAddresses(request *DescribeIPAddressesRequest) (response *DescribeIPAddressesResponse, err error) {
	if request == nil {
		request = NewDescribeIPAddressesRequest()
	}
	response = NewDescribeIPAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetworkPortsRequest() (request *ModifyNetworkPortsRequest) {
	request = &ModifyNetworkPortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyNetworkPorts")
	return
}

func NewModifyNetworkPortsResponse() (response *ModifyNetworkPortsResponse) {
	response = &ModifyNetworkPortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网络设备端口
func (c *Client) ModifyNetworkPorts(request *ModifyNetworkPortsRequest) (response *ModifyNetworkPortsResponse, err error) {
	if request == nil {
		request = NewModifyNetworkPortsRequest()
	}
	response = NewModifyNetworkPortsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDevicesTopoRequest() (request *DescribeNetDevicesTopoRequest) {
	request = &DescribeNetDevicesTopoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevicesTopo")
	return
}

func NewDescribeNetDevicesTopoResponse() (response *DescribeNetDevicesTopoResponse) {
	response = &DescribeNetDevicesTopoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询AZ内网络设备之间物理端口拓扑连接
func (c *Client) DescribeNetDevicesTopo(request *DescribeNetDevicesTopoRequest) (response *DescribeNetDevicesTopoResponse, err error) {
	if request == nil {
		request = NewDescribeNetDevicesTopoRequest()
	}
	response = NewDescribeNetDevicesTopoResponse()
	err = c.Send(request, response)
	return
}

func NewGetWanConfigRequest() (request *GetWanConfigRequest) {
	request = &GetWanConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetWanConfig")
	return
}

func NewGetWanConfigResponse() (response *GetWanConfigResponse) {
	response = &GetWanConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询外网探测配置
func (c *Client) GetWanConfig(request *GetWanConfigRequest) (response *GetWanConfigResponse, err error) {
	if request == nil {
		request = NewGetWanConfigRequest()
	}
	response = NewGetWanConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPSegmentFilterRequest() (request *DescribeIPSegmentFilterRequest) {
	request = &DescribeIPSegmentFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIPSegmentFilter")
	return
}

func NewDescribeIPSegmentFilterResponse() (response *DescribeIPSegmentFilterResponse) {
	response = &DescribeIPSegmentFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取IP网段的过滤值范围
func (c *Client) DescribeIPSegmentFilter(request *DescribeIPSegmentFilterRequest) (response *DescribeIPSegmentFilterResponse, err error) {
	if request == nil {
		request = NewDescribeIPSegmentFilterRequest()
	}
	response = NewDescribeIPSegmentFilterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDevicePortChannelListRequest() (request *DescribeNetDevicePortChannelListRequest) {
	request = &DescribeNetDevicePortChannelListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevicePortChannelList")
	return
}

func NewDescribeNetDevicePortChannelListResponse() (response *DescribeNetDevicePortChannelListResponse) {
	response = &DescribeNetDevicePortChannelListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备端口光模块通道列表对象
func (c *Client) DescribeNetDevicePortChannelList(request *DescribeNetDevicePortChannelListRequest) (response *DescribeNetDevicePortChannelListResponse, err error) {
	if request == nil {
		request = NewDescribeNetDevicePortChannelListRequest()
	}
	response = NewDescribeNetDevicePortChannelListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIPCIDRSegmentsRequest() (request *DeleteIPCIDRSegmentsRequest) {
	request = &DeleteIPCIDRSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteIPCIDRSegments")
	return
}

func NewDeleteIPCIDRSegmentsResponse() (response *DeleteIPCIDRSegmentsResponse) {
	response = &DeleteIPCIDRSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除IPCIDR网段信息
func (c *Client) DeleteIPCIDRSegments(request *DeleteIPCIDRSegmentsRequest) (response *DeleteIPCIDRSegmentsResponse, err error) {
	if request == nil {
		request = NewDeleteIPCIDRSegmentsRequest()
	}
	response = NewDeleteIPCIDRSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIPSegmentsRequest() (request *DescribeIPSegmentsRequest) {
	request = &DescribeIPSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIPSegments")
	return
}

func NewDescribeIPSegmentsResponse() (response *DescribeIPSegmentsResponse) {
	response = &DescribeIPSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取IP网段信息
func (c *Client) DescribeIPSegments(request *DescribeIPSegmentsRequest) (response *DescribeIPSegmentsResponse, err error) {
	if request == nil {
		request = NewDescribeIPSegmentsRequest()
	}
	response = NewDescribeIPSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewSaveTopoSettingRequest() (request *SaveTopoSettingRequest) {
	request = &SaveTopoSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "SaveTopoSetting")
	return
}

func NewSaveTopoSettingResponse() (response *SaveTopoSettingResponse) {
	response = &SaveTopoSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存用户网络拓扑位置配置
func (c *Client) SaveTopoSetting(request *SaveTopoSettingRequest) (response *SaveTopoSettingResponse, err error) {
	if request == nil {
		request = NewSaveTopoSettingRequest()
	}
	response = NewSaveTopoSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecialIdcLineRequest() (request *DescribeSpecialIdcLineRequest) {
	request = &DescribeSpecialIdcLineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeSpecialIdcLine")
	return
}

func NewDescribeSpecialIdcLineResponse() (response *DescribeSpecialIdcLineResponse) {
	response = &DescribeSpecialIdcLineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询IDC专线
func (c *Client) DescribeSpecialIdcLine(request *DescribeSpecialIdcLineRequest) (response *DescribeSpecialIdcLineResponse, err error) {
	if request == nil {
		request = NewDescribeSpecialIdcLineRequest()
	}
	response = NewDescribeSpecialIdcLineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkDevicesRequest() (request *CreateNetworkDevicesRequest) {
	request = &CreateNetworkDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateNetworkDevices")
	return
}

func NewCreateNetworkDevicesResponse() (response *CreateNetworkDevicesResponse) {
	response = &CreateNetworkDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网络设备信息
func (c *Client) CreateNetworkDevices(request *CreateNetworkDevicesRequest) (response *CreateNetworkDevicesResponse, err error) {
	if request == nil {
		request = NewCreateNetworkDevicesRequest()
	}
	response = NewCreateNetworkDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewGetWanDataRequest() (request *GetWanDataRequest) {
	request = &GetWanDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetWanData")
	return
}

func NewGetWanDataResponse() (response *GetWanDataResponse) {
	response = &GetWanDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询外网探测数据
func (c *Client) GetWanData(request *GetWanDataRequest) (response *GetWanDataResponse, err error) {
	if request == nil {
		request = NewGetWanDataRequest()
	}
	response = NewGetWanDataResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRackPositionsRequest() (request *ModifyRackPositionsRequest) {
	request = &ModifyRackPositionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyRackPositions")
	return
}

func NewModifyRackPositionsResponse() (response *ModifyRackPositionsResponse) {
	response = &ModifyRackPositionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机位信息
func (c *Client) ModifyRackPositions(request *ModifyRackPositionsRequest) (response *ModifyRackPositionsResponse, err error) {
	if request == nil {
		request = NewModifyRackPositionsRequest()
	}
	response = NewModifyRackPositionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrapAlarmAbstractRequest() (request *DescribeTrapAlarmAbstractRequest) {
	request = &DescribeTrapAlarmAbstractRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeTrapAlarmAbstract")
	return
}

func NewDescribeTrapAlarmAbstractResponse() (response *DescribeTrapAlarmAbstractResponse) {
	response = &DescribeTrapAlarmAbstractResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询SNMP TRAP概要告警信息
func (c *Client) DescribeTrapAlarmAbstract(request *DescribeTrapAlarmAbstractRequest) (response *DescribeTrapAlarmAbstractResponse, err error) {
	if request == nil {
		request = NewDescribeTrapAlarmAbstractRequest()
	}
	response = NewDescribeTrapAlarmAbstractResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewaysRequest() (request *CreateGatewaysRequest) {
	request = &CreateGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateGateways")
	return
}

func NewCreateGatewaysResponse() (response *CreateGatewaysResponse) {
	response = &CreateGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网关
func (c *Client) CreateGateways(request *CreateGatewaysRequest) (response *CreateGatewaysResponse, err error) {
	if request == nil {
		request = NewCreateGatewaysRequest()
	}
	response = NewCreateGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopologyNetDeviceCountRequest() (request *DescribeTopologyNetDeviceCountRequest) {
	request = &DescribeTopologyNetDeviceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeTopologyNetDeviceCount")
	return
}

func NewDescribeTopologyNetDeviceCountResponse() (response *DescribeTopologyNetDeviceCountResponse) {
	response = &DescribeTopologyNetDeviceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络拓扑内网络设备数目
func (c *Client) DescribeTopologyNetDeviceCount(request *DescribeTopologyNetDeviceCountRequest) (response *DescribeTopologyNetDeviceCountResponse, err error) {
	if request == nil {
		request = NewDescribeTopologyNetDeviceCountRequest()
	}
	response = NewDescribeTopologyNetDeviceCountResponse()
	err = c.Send(request, response)
	return
}

func NewImportSpecialIdcLineFromDcosToNmsRequest() (request *ImportSpecialIdcLineFromDcosToNmsRequest) {
	request = &ImportSpecialIdcLineFromDcosToNmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ImportSpecialIdcLineFromDcosToNms")
	return
}

func NewImportSpecialIdcLineFromDcosToNmsResponse() (response *ImportSpecialIdcLineFromDcosToNmsResponse) {
	response = &ImportSpecialIdcLineFromDcosToNmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步IDC专线信息从CMDB到NMSDB
func (c *Client) ImportSpecialIdcLineFromDcosToNms(request *ImportSpecialIdcLineFromDcosToNmsRequest) (response *ImportSpecialIdcLineFromDcosToNmsResponse, err error) {
	if request == nil {
		request = NewImportSpecialIdcLineFromDcosToNmsRequest()
	}
	response = NewImportSpecialIdcLineFromDcosToNmsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewaysRequest() (request *DeleteGatewaysRequest) {
	request = &DeleteGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteGateways")
	return
}

func NewDeleteGatewaysResponse() (response *DeleteGatewaysResponse) {
	response = &DeleteGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关
func (c *Client) DeleteGateways(request *DeleteGatewaysRequest) (response *DeleteGatewaysResponse, err error) {
	if request == nil {
		request = NewDeleteGatewaysRequest()
	}
	response = NewDeleteGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDevicePortStateRequest() (request *DescribeNetDevicePortStateRequest) {
	request = &DescribeNetDevicePortStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevicePortState")
	return
}

func NewDescribeNetDevicePortStateResponse() (response *DescribeNetDevicePortStateResponse) {
	response = &DescribeNetDevicePortStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备端口状态
func (c *Client) DescribeNetDevicePortState(request *DescribeNetDevicePortStateRequest) (response *DescribeNetDevicePortStateResponse, err error) {
	if request == nil {
		request = NewDescribeNetDevicePortStateRequest()
	}
	response = NewDescribeNetDevicePortStateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRacksRequest() (request *CreateRacksRequest) {
	request = &CreateRacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateRacks")
	return
}

func NewCreateRacksResponse() (response *CreateRacksResponse) {
	response = &CreateRacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机架信息
func (c *Client) CreateRacks(request *CreateRacksRequest) (response *CreateRacksResponse, err error) {
	if request == nil {
		request = NewCreateRacksRequest()
	}
	response = NewCreateRacksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewaysRequest() (request *DescribeGatewaysRequest) {
	request = &DescribeGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeGateways")
	return
}

func NewDescribeGatewaysResponse() (response *DescribeGatewaysResponse) {
	response = &DescribeGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据过滤条件查询网关
func (c *Client) DescribeGateways(request *DescribeGatewaysRequest) (response *DescribeGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeGatewaysRequest()
	}
	response = NewDescribeGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIdcOverviewRequest() (request *DescribeIdcOverviewRequest) {
	request = &DescribeIdcOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIdcOverview")
	return
}

func NewDescribeIdcOverviewResponse() (response *DescribeIdcOverviewResponse) {
	response = &DescribeIdcOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询idc机架机位总览视图
func (c *Client) DescribeIdcOverview(request *DescribeIdcOverviewRequest) (response *DescribeIdcOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeIdcOverviewRequest()
	}
	response = NewDescribeIdcOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateIPAddressesRequest() (request *UpdateIPAddressesRequest) {
	request = &UpdateIPAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "UpdateIPAddresses")
	return
}

func NewUpdateIPAddressesResponse() (response *UpdateIPAddressesResponse) {
	response = &UpdateIPAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新IP地址
func (c *Client) UpdateIPAddresses(request *UpdateIPAddressesRequest) (response *UpdateIPAddressesResponse, err error) {
	if request == nil {
		request = NewUpdateIPAddressesRequest()
	}
	response = NewUpdateIPAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTrapAlarmRequest() (request *DeleteTrapAlarmRequest) {
	request = &DeleteTrapAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteTrapAlarm")
	return
}

func NewDeleteTrapAlarmResponse() (response *DeleteTrapAlarmResponse) {
	response = &DeleteTrapAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除SNMP TRAP告警信息
func (c *Client) DeleteTrapAlarm(request *DeleteTrapAlarmRequest) (response *DeleteTrapAlarmResponse, err error) {
	if request == nil {
		request = NewDeleteTrapAlarmRequest()
	}
	response = NewDeleteTrapAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRackExtendRequest() (request *UpdateRackExtendRequest) {
	request = &UpdateRackExtendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "UpdateRackExtend")
	return
}

func NewUpdateRackExtendResponse() (response *UpdateRackExtendResponse) {
	response = &UpdateRackExtendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机架信息
func (c *Client) UpdateRackExtend(request *UpdateRackExtendRequest) (response *UpdateRackExtendResponse, err error) {
	if request == nil {
		request = NewUpdateRackExtendRequest()
	}
	response = NewUpdateRackExtendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDashboardInfoRequest() (request *DescribeDashboardInfoRequest) {
	request = &DescribeDashboardInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeDashboardInfo")
	return
}

func NewDescribeDashboardInfoResponse() (response *DescribeDashboardInfoResponse) {
	response = &DescribeDashboardInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询概览页
func (c *Client) DescribeDashboardInfo(request *DescribeDashboardInfoRequest) (response *DescribeDashboardInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDashboardInfoRequest()
	}
	response = NewDescribeDashboardInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetDeviceSyslogKeywordRequest() (request *CreateNetDeviceSyslogKeywordRequest) {
	request = &CreateNetDeviceSyslogKeywordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateNetDeviceSyslogKeyword")
	return
}

func NewCreateNetDeviceSyslogKeywordResponse() (response *CreateNetDeviceSyslogKeywordResponse) {
	response = &CreateNetDeviceSyslogKeywordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网络设备系统日志关键字，匹配到关键字后会上报事件
func (c *Client) CreateNetDeviceSyslogKeyword(request *CreateNetDeviceSyslogKeywordRequest) (response *CreateNetDeviceSyslogKeywordResponse, err error) {
	if request == nil {
		request = NewCreateNetDeviceSyslogKeywordRequest()
	}
	response = NewCreateNetDeviceSyslogKeywordResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIPSegmentsRequest() (request *DeleteIPSegmentsRequest) {
	request = &DeleteIPSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteIPSegments")
	return
}

func NewDeleteIPSegmentsResponse() (response *DeleteIPSegmentsResponse) {
	response = &DeleteIPSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除IP网段
func (c *Client) DeleteIPSegments(request *DeleteIPSegmentsRequest) (response *DeleteIPSegmentsResponse, err error) {
	if request == nil {
		request = NewDeleteIPSegmentsRequest()
	}
	response = NewDeleteIPSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDeviceSlotEnumRequest() (request *DescribeNetDeviceSlotEnumRequest) {
	request = &DescribeNetDeviceSlotEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDeviceSlotEnum")
	return
}

func NewDescribeNetDeviceSlotEnumResponse() (response *DescribeNetDeviceSlotEnumResponse) {
	response = &DescribeNetDeviceSlotEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备单板枚举信息
func (c *Client) DescribeNetDeviceSlotEnum(request *DescribeNetDeviceSlotEnumRequest) (response *DescribeNetDeviceSlotEnumResponse, err error) {
	if request == nil {
		request = NewDescribeNetDeviceSlotEnumRequest()
	}
	response = NewDescribeNetDeviceSlotEnumResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFieldsEnumRequest() (request *DescribeFieldsEnumRequest) {
	request = &DescribeFieldsEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeFieldsEnum")
	return
}

func NewDescribeFieldsEnumResponse() (response *DescribeFieldsEnumResponse) {
	response = &DescribeFieldsEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备各字段枚举值
func (c *Client) DescribeFieldsEnum(request *DescribeFieldsEnumRequest) (response *DescribeFieldsEnumResponse, err error) {
	if request == nil {
		request = NewDescribeFieldsEnumRequest()
	}
	response = NewDescribeFieldsEnumResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecialIdcLineTopoRequest() (request *DescribeSpecialIdcLineTopoRequest) {
	request = &DescribeSpecialIdcLineTopoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeSpecialIdcLineTopo")
	return
}

func NewDescribeSpecialIdcLineTopoResponse() (response *DescribeSpecialIdcLineTopoResponse) {
	response = &DescribeSpecialIdcLineTopoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Region内AZ之间的专线拓扑连接关系。
func (c *Client) DescribeSpecialIdcLineTopo(request *DescribeSpecialIdcLineTopoRequest) (response *DescribeSpecialIdcLineTopoResponse, err error) {
	if request == nil {
		request = NewDescribeSpecialIdcLineTopoRequest()
	}
	response = NewDescribeSpecialIdcLineTopoResponse()
	err = c.Send(request, response)
	return
}

func NewDciCapacityRequest() (request *DciCapacityRequest) {
	request = &DciCapacityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DciCapacity")
	return
}

func NewDciCapacityResponse() (response *DciCapacityResponse) {
	response = &DciCapacityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DCI容量查询列表
func (c *Client) DciCapacity(request *DciCapacityRequest) (response *DciCapacityResponse, err error) {
	if request == nil {
		request = NewDciCapacityRequest()
	}
	response = NewDciCapacityResponse()
	err = c.Send(request, response)
	return
}

func NewDescirbeGlobalNetDeviceSyslogsRequest() (request *DescirbeGlobalNetDeviceSyslogsRequest) {
	request = &DescirbeGlobalNetDeviceSyslogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescirbeGlobalNetDeviceSyslogs")
	return
}

func NewDescirbeGlobalNetDeviceSyslogsResponse() (response *DescirbeGlobalNetDeviceSyslogsResponse) {
	response = &DescirbeGlobalNetDeviceSyslogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全局的网络设备日志信息
func (c *Client) DescirbeGlobalNetDeviceSyslogs(request *DescirbeGlobalNetDeviceSyslogsRequest) (response *DescirbeGlobalNetDeviceSyslogsResponse, err error) {
	if request == nil {
		request = NewDescirbeGlobalNetDeviceSyslogsRequest()
	}
	response = NewDescirbeGlobalNetDeviceSyslogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkDevicesExRequest() (request *DescribeNetworkDevicesExRequest) {
	request = &DescribeNetworkDevicesExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetworkDevicesEx")
	return
}

func NewDescribeNetworkDevicesExResponse() (response *DescribeNetworkDevicesExResponse) {
	response = &DescribeNetworkDevicesExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备信息
func (c *Client) DescribeNetworkDevicesEx(request *DescribeNetworkDevicesExRequest) (response *DescribeNetworkDevicesExResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkDevicesExRequest()
	}
	response = NewDescribeNetworkDevicesExResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayClustersRequest() (request *UpdateGatewayClustersRequest) {
	request = &UpdateGatewayClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "UpdateGatewayClusters")
	return
}

func NewUpdateGatewayClustersResponse() (response *UpdateGatewayClustersResponse) {
	response = &UpdateGatewayClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新网关集群信息
func (c *Client) UpdateGatewayClusters(request *UpdateGatewayClustersRequest) (response *UpdateGatewayClustersResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayClustersRequest()
	}
	response = NewUpdateGatewayClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRacksRequest() (request *DescribeRacksRequest) {
	request = &DescribeRacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeRacks")
	return
}

func NewDescribeRacksResponse() (response *DescribeRacksResponse) {
	response = &DescribeRacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机架信息
func (c *Client) DescribeRacks(request *DescribeRacksRequest) (response *DescribeRacksResponse, err error) {
	if request == nil {
		request = NewDescribeRacksRequest()
	}
	response = NewDescribeRacksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTemplateConfigRequest() (request *DeleteTemplateConfigRequest) {
	request = &DeleteTemplateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteTemplateConfig")
	return
}

func NewDeleteTemplateConfigResponse() (response *DeleteTemplateConfigResponse) {
	response = &DeleteTemplateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络设备配置模板绑定
func (c *Client) DeleteTemplateConfig(request *DeleteTemplateConfigRequest) (response *DeleteTemplateConfigResponse, err error) {
	if request == nil {
		request = NewDeleteTemplateConfigRequest()
	}
	response = NewDeleteTemplateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRackPositionsRequest() (request *DescribeRackPositionsRequest) {
	request = &DescribeRackPositionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeRackPositions")
	return
}

func NewDescribeRackPositionsResponse() (response *DescribeRackPositionsResponse) {
	response = &DescribeRackPositionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机位信息
func (c *Client) DescribeRackPositions(request *DescribeRackPositionsRequest) (response *DescribeRackPositionsResponse, err error) {
	if request == nil {
		request = NewDescribeRackPositionsRequest()
	}
	response = NewDescribeRackPositionsResponse()
	err = c.Send(request, response)
	return
}

func NewWithdrawNetworkDevicesRequest() (request *WithdrawNetworkDevicesRequest) {
	request = &WithdrawNetworkDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "WithdrawNetworkDevices")
	return
}

func NewWithdrawNetworkDevicesResponse() (response *WithdrawNetworkDevicesResponse) {
	response = &WithdrawNetworkDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下架网络设备
func (c *Client) WithdrawNetworkDevices(request *WithdrawNetworkDevicesRequest) (response *WithdrawNetworkDevicesResponse, err error) {
	if request == nil {
		request = NewWithdrawNetworkDevicesRequest()
	}
	response = NewWithdrawNetworkDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkRoleDictRequest() (request *DeleteNetworkRoleDictRequest) {
	request = &DeleteNetworkRoleDictRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteNetworkRoleDict")
	return
}

func NewDeleteNetworkRoleDictResponse() (response *DeleteNetworkRoleDictResponse) {
	response = &DeleteNetworkRoleDictResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络设备角色中英文字典
func (c *Client) DeleteNetworkRoleDict(request *DeleteNetworkRoleDictRequest) (response *DeleteNetworkRoleDictResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkRoleDictRequest()
	}
	response = NewDeleteNetworkRoleDictResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDeviceSlotListRequest() (request *DescribeNetDeviceSlotListRequest) {
	request = &DescribeNetDeviceSlotListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDeviceSlotList")
	return
}

func NewDescribeNetDeviceSlotListResponse() (response *DescribeNetDeviceSlotListResponse) {
	response = &DescribeNetDeviceSlotListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备单板列表信息
func (c *Client) DescribeNetDeviceSlotList(request *DescribeNetDeviceSlotListRequest) (response *DescribeNetDeviceSlotListResponse, err error) {
	if request == nil {
		request = NewDescribeNetDeviceSlotListRequest()
	}
	response = NewDescribeNetDeviceSlotListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetdeviceFieldsEnumRequest() (request *DescribeNetdeviceFieldsEnumRequest) {
	request = &DescribeNetdeviceFieldsEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetdeviceFieldsEnum")
	return
}

func NewDescribeNetdeviceFieldsEnumResponse() (response *DescribeNetdeviceFieldsEnumResponse) {
	response = &DescribeNetdeviceFieldsEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备字段枚举
func (c *Client) DescribeNetdeviceFieldsEnum(request *DescribeNetdeviceFieldsEnumRequest) (response *DescribeNetdeviceFieldsEnumResponse, err error) {
	if request == nil {
		request = NewDescribeNetdeviceFieldsEnumRequest()
	}
	response = NewDescribeNetdeviceFieldsEnumResponse()
	err = c.Send(request, response)
	return
}

func NewImportNetDeviceSyslogKeywordRequest() (request *ImportNetDeviceSyslogKeywordRequest) {
	request = &ImportNetDeviceSyslogKeywordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ImportNetDeviceSyslogKeyword")
	return
}

func NewImportNetDeviceSyslogKeywordResponse() (response *ImportNetDeviceSyslogKeywordResponse) {
	response = &ImportNetDeviceSyslogKeywordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入网络设备系统日志关键字
func (c *Client) ImportNetDeviceSyslogKeyword(request *ImportNetDeviceSyslogKeywordRequest) (response *ImportNetDeviceSyslogKeywordResponse, err error) {
	if request == nil {
		request = NewImportNetDeviceSyslogKeywordRequest()
	}
	response = NewImportNetDeviceSyslogKeywordResponse()
	err = c.Send(request, response)
	return
}

func NewImportIPAddressesRequest() (request *ImportIPAddressesRequest) {
	request = &ImportIPAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ImportIPAddresses")
	return
}

func NewImportIPAddressesResponse() (response *ImportIPAddressesResponse) {
	response = &ImportIPAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入IP
func (c *Client) ImportIPAddresses(request *ImportIPAddressesRequest) (response *ImportIPAddressesResponse, err error) {
	if request == nil {
		request = NewImportIPAddressesRequest()
	}
	response = NewImportIPAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCPUUsageTop5Request() (request *DescribeCPUUsageTop5Request) {
	request = &DescribeCPUUsageTop5Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeCPUUsageTop5")
	return
}

func NewDescribeCPUUsageTop5Response() (response *DescribeCPUUsageTop5Response) {
	response = &DescribeCPUUsageTop5Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询CPU利用率Top5数据
func (c *Client) DescribeCPUUsageTop5(request *DescribeCPUUsageTop5Request) (response *DescribeCPUUsageTop5Response, err error) {
	if request == nil {
		request = NewDescribeCPUUsageTop5Request()
	}
	response = NewDescribeCPUUsageTop5Response()
	err = c.Send(request, response)
	return
}

func NewUpdateIPCidrSegmentsRequest() (request *UpdateIPCidrSegmentsRequest) {
	request = &UpdateIPCidrSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "UpdateIPCidrSegments")
	return
}

func NewUpdateIPCidrSegmentsResponse() (response *UpdateIPCidrSegmentsResponse) {
	response = &UpdateIPCidrSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改IPCIDR网段白名单
func (c *Client) UpdateIPCidrSegments(request *UpdateIPCidrSegmentsRequest) (response *UpdateIPCidrSegmentsResponse, err error) {
	if request == nil {
		request = NewUpdateIPCidrSegmentsRequest()
	}
	response = NewUpdateIPCidrSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRacksOldRequest() (request *CreateRacksOldRequest) {
	request = &CreateRacksOldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateRacksOld")
	return
}

func NewCreateRacksOldResponse() (response *CreateRacksOldResponse) {
	response = &CreateRacksOldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机架信息
func (c *Client) CreateRacksOld(request *CreateRacksOldRequest) (response *CreateRacksOldResponse, err error) {
	if request == nil {
		request = NewCreateRacksOldRequest()
	}
	response = NewCreateRacksOldResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZoneRequest() (request *DescribeZoneRequest) {
	request = &DescribeZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeZone")
	return
}

func NewDescribeZoneResponse() (response *DescribeZoneResponse) {
	response = &DescribeZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询地域内可用区信息
func (c *Client) DescribeZone(request *DescribeZoneRequest) (response *DescribeZoneResponse, err error) {
	if request == nil {
		request = NewDescribeZoneRequest()
	}
	response = NewDescribeZoneResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkRoleDictRequest() (request *CreateNetworkRoleDictRequest) {
	request = &CreateNetworkRoleDictRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateNetworkRoleDict")
	return
}

func NewCreateNetworkRoleDictResponse() (response *CreateNetworkRoleDictResponse) {
	response = &CreateNetworkRoleDictResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加网络设备角色中英文字典
func (c *Client) CreateNetworkRoleDict(request *CreateNetworkRoleDictRequest) (response *CreateNetworkRoleDictResponse, err error) {
	if request == nil {
		request = NewCreateNetworkRoleDictRequest()
	}
	response = NewCreateNetworkRoleDictResponse()
	err = c.Send(request, response)
	return
}

func NewGetDetailPingDataRequest() (request *GetDetailPingDataRequest) {
	request = &GetDetailPingDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetDetailPingData")
	return
}

func NewGetDetailPingDataResponse() (response *GetDetailPingDataResponse) {
	response = &GetDetailPingDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询探测明细数据
func (c *Client) GetDetailPingData(request *GetDetailPingDataRequest) (response *GetDetailPingDataResponse, err error) {
	if request == nil {
		request = NewGetDetailPingDataRequest()
	}
	response = NewGetDetailPingDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDevicePortEnumRequest() (request *DescribeNetDevicePortEnumRequest) {
	request = &DescribeNetDevicePortEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevicePortEnum")
	return
}

func NewDescribeNetDevicePortEnumResponse() (response *DescribeNetDevicePortEnumResponse) {
	response = &DescribeNetDevicePortEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备端口枚举对象
func (c *Client) DescribeNetDevicePortEnum(request *DescribeNetDevicePortEnumRequest) (response *DescribeNetDevicePortEnumResponse, err error) {
	if request == nil {
		request = NewDescribeNetDevicePortEnumRequest()
	}
	response = NewDescribeNetDevicePortEnumResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewaysRequest() (request *UpdateGatewaysRequest) {
	request = &UpdateGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "UpdateGateways")
	return
}

func NewUpdateGatewaysResponse() (response *UpdateGatewaysResponse) {
	response = &UpdateGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新网关相关信息
func (c *Client) UpdateGateways(request *UpdateGatewaysRequest) (response *UpdateGatewaysResponse, err error) {
	if request == nil {
		request = NewUpdateGatewaysRequest()
	}
	response = NewUpdateGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpecialIdcLineMetricRequest() (request *DescribeSpecialIdcLineMetricRequest) {
	request = &DescribeSpecialIdcLineMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeSpecialIdcLineMetric")
	return
}

func NewDescribeSpecialIdcLineMetricResponse() (response *DescribeSpecialIdcLineMetricResponse) {
	response = &DescribeSpecialIdcLineMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询AZ内IDC专线的指标数据
func (c *Client) DescribeSpecialIdcLineMetric(request *DescribeSpecialIdcLineMetricRequest) (response *DescribeSpecialIdcLineMetricResponse, err error) {
	if request == nil {
		request = NewDescribeSpecialIdcLineMetricRequest()
	}
	response = NewDescribeSpecialIdcLineMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkDevicesRequest() (request *DescribeNetworkDevicesRequest) {
	request = &DescribeNetworkDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetworkDevices")
	return
}

func NewDescribeNetworkDevicesResponse() (response *DescribeNetworkDevicesResponse) {
	response = &DescribeNetworkDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备信息
func (c *Client) DescribeNetworkDevices(request *DescribeNetworkDevicesRequest) (response *DescribeNetworkDevicesResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkDevicesRequest()
	}
	response = NewDescribeNetworkDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIdcDirectConnectTopoRequest() (request *DescribeIdcDirectConnectTopoRequest) {
	request = &DescribeIdcDirectConnectTopoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIdcDirectConnectTopo")
	return
}

func NewDescribeIdcDirectConnectTopoResponse() (response *DescribeIdcDirectConnectTopoResponse) {
	response = &DescribeIdcDirectConnectTopoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询物理专线的拓扑连接关系
func (c *Client) DescribeIdcDirectConnectTopo(request *DescribeIdcDirectConnectTopoRequest) (response *DescribeIdcDirectConnectTopoResponse, err error) {
	if request == nil {
		request = NewDescribeIdcDirectConnectTopoRequest()
	}
	response = NewDescribeIdcDirectConnectTopoResponse()
	err = c.Send(request, response)
	return
}

func NewGetDciConfigRequest() (request *GetDciConfigRequest) {
	request = &GetDciConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetDciConfig")
	return
}

func NewGetDciConfigResponse() (response *GetDciConfigResponse) {
	response = &GetDciConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取DCI探测配置
func (c *Client) GetDciConfig(request *GetDciConfigRequest) (response *GetDciConfigResponse, err error) {
	if request == nil {
		request = NewGetDciConfigRequest()
	}
	response = NewGetDciConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayClustersRequest() (request *DescribeGatewayClustersRequest) {
	request = &DescribeGatewayClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeGatewayClusters")
	return
}

func NewDescribeGatewayClustersResponse() (response *DescribeGatewayClustersResponse) {
	response = &DescribeGatewayClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据输入条件获取网关集群列表
func (c *Client) DescribeGatewayClusters(request *DescribeGatewayClustersRequest) (response *DescribeGatewayClustersResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayClustersRequest()
	}
	response = NewDescribeGatewayClustersResponse()
	err = c.Send(request, response)
	return
}

func NewGetLanConfigRequest() (request *GetLanConfigRequest) {
	request = &GetLanConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetLanConfig")
	return
}

func NewGetLanConfigResponse() (response *GetLanConfigResponse) {
	response = &GetLanConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内网探测配置
func (c *Client) GetLanConfig(request *GetLanConfigRequest) (response *GetLanConfigResponse, err error) {
	if request == nil {
		request = NewGetLanConfigRequest()
	}
	response = NewGetLanConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIdcExportLineMetricRequest() (request *DescribeIdcExportLineMetricRequest) {
	request = &DescribeIdcExportLineMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIdcExportLineMetric")
	return
}

func NewDescribeIdcExportLineMetricResponse() (response *DescribeIdcExportLineMetricResponse) {
	response = &DescribeIdcExportLineMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询AZ内IDC出口的指标数据
func (c *Client) DescribeIdcExportLineMetric(request *DescribeIdcExportLineMetricRequest) (response *DescribeIdcExportLineMetricResponse, err error) {
	if request == nil {
		request = NewDescribeIdcExportLineMetricRequest()
	}
	response = NewDescribeIdcExportLineMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDciSpeLineListRequest() (request *DciSpeLineListRequest) {
	request = &DciSpeLineListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DciSpeLineList")
	return
}

func NewDciSpeLineListResponse() (response *DciSpeLineListResponse) {
	response = &DciSpeLineListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DCI容量查询专线列表
func (c *Client) DciSpeLineList(request *DciSpeLineListRequest) (response *DciSpeLineListResponse, err error) {
	if request == nil {
		request = NewDciSpeLineListRequest()
	}
	response = NewDciSpeLineListResponse()
	err = c.Send(request, response)
	return
}

func NewGetWanExitStatusStatRequest() (request *GetWanExitStatusStatRequest) {
	request = &GetWanExitStatusStatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetWanExitStatusStat")
	return
}

func NewGetWanExitStatusStatResponse() (response *GetWanExitStatusStatResponse) {
	response = &GetWanExitStatusStatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询IDC出口状态
func (c *Client) GetWanExitStatusStat(request *GetWanExitStatusStatRequest) (response *GetWanExitStatusStatResponse, err error) {
	if request == nil {
		request = NewGetWanExitStatusStatRequest()
	}
	response = NewGetWanExitStatusStatResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRackPositionsRequest() (request *CreateRackPositionsRequest) {
	request = &CreateRackPositionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateRackPositions")
	return
}

func NewCreateRackPositionsResponse() (response *CreateRackPositionsResponse) {
	response = &CreateRackPositionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机位信息
func (c *Client) CreateRackPositions(request *CreateRackPositionsRequest) (response *CreateRackPositionsResponse, err error) {
	if request == nil {
		request = NewCreateRackPositionsRequest()
	}
	response = NewCreateRackPositionsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateNetDeviceSyslogKeywordRequest() (request *UpdateNetDeviceSyslogKeywordRequest) {
	request = &UpdateNetDeviceSyslogKeywordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "UpdateNetDeviceSyslogKeyword")
	return
}

func NewUpdateNetDeviceSyslogKeywordResponse() (response *UpdateNetDeviceSyslogKeywordResponse) {
	response = &UpdateNetDeviceSyslogKeywordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新网络设备系统日志关键字
func (c *Client) UpdateNetDeviceSyslogKeyword(request *UpdateNetDeviceSyslogKeywordRequest) (response *UpdateNetDeviceSyslogKeywordResponse, err error) {
	if request == nil {
		request = NewUpdateNetDeviceSyslogKeywordRequest()
	}
	response = NewUpdateNetDeviceSyslogKeywordResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNetDeviceConfigBaseLineRequest() (request *ModifyNetDeviceConfigBaseLineRequest) {
	request = &ModifyNetDeviceConfigBaseLineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyNetDeviceConfigBaseLine")
	return
}

func NewModifyNetDeviceConfigBaseLineResponse() (response *ModifyNetDeviceConfigBaseLineResponse) {
	response = &ModifyNetDeviceConfigBaseLineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置网络设备配置基线
func (c *Client) ModifyNetDeviceConfigBaseLine(request *ModifyNetDeviceConfigBaseLineRequest) (response *ModifyNetDeviceConfigBaseLineResponse, err error) {
	if request == nil {
		request = NewModifyNetDeviceConfigBaseLineRequest()
	}
	response = NewModifyNetDeviceConfigBaseLineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServersTopoRequest() (request *DescribeServersTopoRequest) {
	request = &DescribeServersTopoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeServersTopo")
	return
}

func NewDescribeServersTopoResponse() (response *DescribeServersTopoResponse) {
	response = &DescribeServersTopoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备和服务器之间拓扑连接关系，展示连接设备系统名称，物理端口。
func (c *Client) DescribeServersTopo(request *DescribeServersTopoRequest) (response *DescribeServersTopoResponse, err error) {
	if request == nil {
		request = NewDescribeServersTopoRequest()
	}
	response = NewDescribeServersTopoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceAlarmRequest() (request *DescribeDeviceAlarmRequest) {
	request = &DescribeDeviceAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeDeviceAlarm")
	return
}

func NewDescribeDeviceAlarmResponse() (response *DescribeDeviceAlarmResponse) {
	response = &DescribeDeviceAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警设备
func (c *Client) DescribeDeviceAlarm(request *DescribeDeviceAlarmRequest) (response *DescribeDeviceAlarmResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceAlarmRequest()
	}
	response = NewDescribeDeviceAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMemoryUsageTop5Request() (request *DescribeMemoryUsageTop5Request) {
	request = &DescribeMemoryUsageTop5Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeMemoryUsageTop5")
	return
}

func NewDescribeMemoryUsageTop5Response() (response *DescribeMemoryUsageTop5Response) {
	response = &DescribeMemoryUsageTop5Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内存利用率Top5
func (c *Client) DescribeMemoryUsageTop5(request *DescribeMemoryUsageTop5Request) (response *DescribeMemoryUsageTop5Response, err error) {
	if request == nil {
		request = NewDescribeMemoryUsageTop5Request()
	}
	response = NewDescribeMemoryUsageTop5Response()
	err = c.Send(request, response)
	return
}

func NewDescribeBandWidthUsageTop5Request() (request *DescribeBandWidthUsageTop5Request) {
	request = &DescribeBandWidthUsageTop5Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeBandWidthUsageTop5")
	return
}

func NewDescribeBandWidthUsageTop5Response() (response *DescribeBandWidthUsageTop5Response) {
	response = &DescribeBandWidthUsageTop5Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询带宽利用率top5
func (c *Client) DescribeBandWidthUsageTop5(request *DescribeBandWidthUsageTop5Request) (response *DescribeBandWidthUsageTop5Response, err error) {
	if request == nil {
		request = NewDescribeBandWidthUsageTop5Request()
	}
	response = NewDescribeBandWidthUsageTop5Response()
	err = c.Send(request, response)
	return
}

func NewDeleteTopoSettingRequest() (request *DeleteTopoSettingRequest) {
	request = &DeleteTopoSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteTopoSetting")
	return
}

func NewDeleteTopoSettingResponse() (response *DeleteTopoSettingResponse) {
	response = &DeleteTopoSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户网络拓扑位置配置
func (c *Client) DeleteTopoSetting(request *DeleteTopoSettingRequest) (response *DeleteTopoSettingResponse, err error) {
	if request == nil {
		request = NewDeleteTopoSettingRequest()
	}
	response = NewDeleteTopoSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDevicePortChannelEnumRequest() (request *DescribeNetDevicePortChannelEnumRequest) {
	request = &DescribeNetDevicePortChannelEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevicePortChannelEnum")
	return
}

func NewDescribeNetDevicePortChannelEnumResponse() (response *DescribeNetDevicePortChannelEnumResponse) {
	response = &DescribeNetDevicePortChannelEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备端口光模块通道枚举对象
func (c *Client) DescribeNetDevicePortChannelEnum(request *DescribeNetDevicePortChannelEnumRequest) (response *DescribeNetDevicePortChannelEnumResponse, err error) {
	if request == nil {
		request = NewDescribeNetDevicePortChannelEnumRequest()
	}
	response = NewDescribeNetDevicePortChannelEnumResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIPAddressesRequest() (request *CreateIPAddressesRequest) {
	request = &CreateIPAddressesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateIPAddresses")
	return
}

func NewCreateIPAddressesResponse() (response *CreateIPAddressesResponse) {
	response = &CreateIPAddressesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 申请IP地址
func (c *Client) CreateIPAddresses(request *CreateIPAddressesRequest) (response *CreateIPAddressesResponse, err error) {
	if request == nil {
		request = NewCreateIPAddressesRequest()
	}
	response = NewCreateIPAddressesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationLogRequest() (request *DescribeOperationLogRequest) {
	request = &DescribeOperationLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeOperationLog")
	return
}

func NewDescribeOperationLogResponse() (response *DescribeOperationLogResponse) {
	response = &DescribeOperationLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作日志记录
func (c *Client) DescribeOperationLog(request *DescribeOperationLogRequest) (response *DescribeOperationLogResponse, err error) {
	if request == nil {
		request = NewDescribeOperationLogRequest()
	}
	response = NewDescribeOperationLogResponse()
	err = c.Send(request, response)
	return
}

func NewGetCommonServicePingDataRequest() (request *GetCommonServicePingDataRequest) {
	request = &GetCommonServicePingDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetCommonServicePingData")
	return
}

func NewGetCommonServicePingDataResponse() (response *GetCommonServicePingDataResponse) {
	response = &GetCommonServicePingDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共网络服务探测数据
func (c *Client) GetCommonServicePingData(request *GetCommonServicePingDataRequest) (response *GetCommonServicePingDataResponse, err error) {
	if request == nil {
		request = NewGetCommonServicePingDataRequest()
	}
	response = NewGetCommonServicePingDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescirbeNetDeviceSyslogsRequest() (request *DescirbeNetDeviceSyslogsRequest) {
	request = &DescirbeNetDeviceSyslogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescirbeNetDeviceSyslogs")
	return
}

func NewDescirbeNetDeviceSyslogsResponse() (response *DescirbeNetDeviceSyslogsResponse) {
	response = &DescirbeNetDeviceSyslogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网络设备的系统日志
func (c *Client) DescirbeNetDeviceSyslogs(request *DescirbeNetDeviceSyslogsRequest) (response *DescirbeNetDeviceSyslogsResponse, err error) {
	if request == nil {
		request = NewDescirbeNetDeviceSyslogsRequest()
	}
	response = NewDescirbeNetDeviceSyslogsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIPSegmentsRequest() (request *CreateIPSegmentsRequest) {
	request = &CreateIPSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateIPSegments")
	return
}

func NewCreateIPSegmentsResponse() (response *CreateIPSegmentsResponse) {
	response = &CreateIPSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加IP网段
func (c *Client) CreateIPSegments(request *CreateIPSegmentsRequest) (response *CreateIPSegmentsResponse, err error) {
	if request == nil {
		request = NewCreateIPSegmentsRequest()
	}
	response = NewCreateIPSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRackPositionsRequest() (request *DeleteRackPositionsRequest) {
	request = &DeleteRackPositionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteRackPositions")
	return
}

func NewDeleteRackPositionsResponse() (response *DeleteRackPositionsResponse) {
	response = &DeleteRackPositionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机位
func (c *Client) DeleteRackPositions(request *DeleteRackPositionsRequest) (response *DeleteRackPositionsResponse, err error) {
	if request == nil {
		request = NewDeleteRackPositionsRequest()
	}
	response = NewDeleteRackPositionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTrapAlarmLevelRequest() (request *ModifyTrapAlarmLevelRequest) {
	request = &ModifyTrapAlarmLevelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyTrapAlarmLevel")
	return
}

func NewModifyTrapAlarmLevelResponse() (response *ModifyTrapAlarmLevelResponse) {
	response = &ModifyTrapAlarmLevelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改SNMP TRAP告警等级
func (c *Client) ModifyTrapAlarmLevel(request *ModifyTrapAlarmLevelRequest) (response *ModifyTrapAlarmLevelResponse, err error) {
	if request == nil {
		request = NewModifyTrapAlarmLevelRequest()
	}
	response = NewModifyTrapAlarmLevelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetDeviceSyslogKeywordRequest() (request *DeleteNetDeviceSyslogKeywordRequest) {
	request = &DeleteNetDeviceSyslogKeywordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteNetDeviceSyslogKeyword")
	return
}

func NewDeleteNetDeviceSyslogKeywordResponse() (response *DeleteNetDeviceSyslogKeywordResponse) {
	response = &DeleteNetDeviceSyslogKeywordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络设备系统日志关键字
func (c *Client) DeleteNetDeviceSyslogKeyword(request *DeleteNetDeviceSyslogKeywordRequest) (response *DeleteNetDeviceSyslogKeywordResponse, err error) {
	if request == nil {
		request = NewDeleteNetDeviceSyslogKeywordRequest()
	}
	response = NewDeleteNetDeviceSyslogKeywordResponse()
	err = c.Send(request, response)
	return
}

func NewGetGwDataRequest() (request *GetGwDataRequest) {
	request = &GetGwDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetGwData")
	return
}

func NewGetGwDataResponse() (response *GetGwDataResponse) {
	response = &GetGwDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关探测数据
func (c *Client) GetGwData(request *GetGwDataRequest) (response *GetGwDataResponse, err error) {
	if request == nil {
		request = NewGetGwDataRequest()
	}
	response = NewGetGwDataResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRackExtendRequest() (request *DeleteRackExtendRequest) {
	request = &DeleteRackExtendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteRackExtend")
	return
}

func NewDeleteRackExtendResponse() (response *DeleteRackExtendResponse) {
	response = &DeleteRackExtendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机柜
func (c *Client) DeleteRackExtend(request *DeleteRackExtendRequest) (response *DeleteRackExtendResponse, err error) {
	if request == nil {
		request = NewDeleteRackExtendRequest()
	}
	response = NewDeleteRackExtendResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNetworkPortsRequest() (request *CreateNetworkPortsRequest) {
	request = &CreateNetworkPortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateNetworkPorts")
	return
}

func NewCreateNetworkPortsResponse() (response *CreateNetworkPortsResponse) {
	response = &CreateNetworkPortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建网络设备端口
func (c *Client) CreateNetworkPorts(request *CreateNetworkPortsRequest) (response *CreateNetworkPortsResponse, err error) {
	if request == nil {
		request = NewCreateNetworkPortsRequest()
	}
	response = NewCreateNetworkPortsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDevicePortListRequest() (request *DescribeNetDevicePortListRequest) {
	request = &DescribeNetDevicePortListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevicePortList")
	return
}

func NewDescribeNetDevicePortListResponse() (response *DescribeNetDevicePortListResponse) {
	response = &DescribeNetDevicePortListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备端口列表对象
func (c *Client) DescribeNetDevicePortList(request *DescribeNetDevicePortListRequest) (response *DescribeNetDevicePortListResponse, err error) {
	if request == nil {
		request = NewDescribeNetDevicePortListRequest()
	}
	response = NewDescribeNetDevicePortListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDevicesTopoExRequest() (request *DescribeNetDevicesTopoExRequest) {
	request = &DescribeNetDevicesTopoExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDevicesTopoEx")
	return
}

func NewDescribeNetDevicesTopoExResponse() (response *DescribeNetDevicesTopoExResponse) {
	response = &DescribeNetDevicesTopoExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询AZ内网络设备之间物理端口拓扑连接
func (c *Client) DescribeNetDevicesTopoEx(request *DescribeNetDevicesTopoExRequest) (response *DescribeNetDevicesTopoExResponse, err error) {
	if request == nil {
		request = NewDescribeNetDevicesTopoExRequest()
	}
	response = NewDescribeNetDevicesTopoExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkPortsRequest() (request *DescribeNetworkPortsRequest) {
	request = &DescribeNetworkPortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetworkPorts")
	return
}

func NewDescribeNetworkPortsResponse() (response *DescribeNetworkPortsResponse) {
	response = &DescribeNetworkPortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备端口信息
func (c *Client) DescribeNetworkPorts(request *DescribeNetworkPortsRequest) (response *DescribeNetworkPortsResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkPortsRequest()
	}
	response = NewDescribeNetworkPortsResponse()
	err = c.Send(request, response)
	return
}

func NewImportIdcExportLineFromDcosToNmsRequest() (request *ImportIdcExportLineFromDcosToNmsRequest) {
	request = &ImportIdcExportLineFromDcosToNmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ImportIdcExportLineFromDcosToNms")
	return
}

func NewImportIdcExportLineFromDcosToNmsResponse() (response *ImportIdcExportLineFromDcosToNmsResponse) {
	response = &ImportIdcExportLineFromDcosToNmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步IDC专线出口信息从CMDB到NMSDB
func (c *Client) ImportIdcExportLineFromDcosToNms(request *ImportIdcExportLineFromDcosToNmsRequest) (response *ImportIdcExportLineFromDcosToNmsResponse, err error) {
	if request == nil {
		request = NewImportIdcExportLineFromDcosToNmsRequest()
	}
	response = NewImportIdcExportLineFromDcosToNmsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetworkPortSortRequest() (request *DescribeNetworkPortSortRequest) {
	request = &DescribeNetworkPortSortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetworkPortSort")
	return
}

func NewDescribeNetworkPortSortResponse() (response *DescribeNetworkPortSortResponse) {
	response = &DescribeNetworkPortSortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网络设备端口排序
func (c *Client) DescribeNetworkPortSort(request *DescribeNetworkPortSortRequest) (response *DescribeNetworkPortSortResponse, err error) {
	if request == nil {
		request = NewDescribeNetworkPortSortRequest()
	}
	response = NewDescribeNetworkPortSortResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSpecialIdcLineRequest() (request *CreateSpecialIdcLineRequest) {
	request = &CreateSpecialIdcLineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateSpecialIdcLine")
	return
}

func NewCreateSpecialIdcLineResponse() (response *CreateSpecialIdcLineResponse) {
	response = &CreateSpecialIdcLineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加IDC专线
func (c *Client) CreateSpecialIdcLine(request *CreateSpecialIdcLineRequest) (response *CreateSpecialIdcLineResponse, err error) {
	if request == nil {
		request = NewCreateSpecialIdcLineRequest()
	}
	response = NewCreateSpecialIdcLineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateConfigRequest() (request *CreateTemplateConfigRequest) {
	request = &CreateTemplateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateTemplateConfig")
	return
}

func NewCreateTemplateConfigResponse() (response *CreateTemplateConfigResponse) {
	response = &CreateTemplateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加网络设备配置模板绑定
func (c *Client) CreateTemplateConfig(request *CreateTemplateConfigRequest) (response *CreateTemplateConfigResponse, err error) {
	if request == nil {
		request = NewCreateTemplateConfigRequest()
	}
	response = NewCreateTemplateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIPCIDRSegmentsRequest() (request *CreateIPCIDRSegmentsRequest) {
	request = &CreateIPCIDRSegmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateIPCIDRSegments")
	return
}

func NewCreateIPCIDRSegmentsResponse() (response *CreateIPCIDRSegmentsResponse) {
	response = &CreateIPCIDRSegmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建IPCIDR网段
func (c *Client) CreateIPCIDRSegments(request *CreateIPCIDRSegmentsRequest) (response *CreateIPCIDRSegmentsResponse, err error) {
	if request == nil {
		request = NewCreateIPCIDRSegmentsRequest()
	}
	response = NewCreateIPCIDRSegmentsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetDeviceConfigsRequest() (request *DescribeNetDeviceConfigsRequest) {
	request = &DescribeNetDeviceConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeNetDeviceConfigs")
	return
}

func NewDescribeNetDeviceConfigsResponse() (response *DescribeNetDeviceConfigsResponse) {
	response = &DescribeNetDeviceConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网络设备配置信息
func (c *Client) DescribeNetDeviceConfigs(request *DescribeNetDeviceConfigsRequest) (response *DescribeNetDeviceConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeNetDeviceConfigsRequest()
	}
	response = NewDescribeNetDeviceConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrapAlarmRequest() (request *DescribeTrapAlarmRequest) {
	request = &DescribeTrapAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeTrapAlarm")
	return
}

func NewDescribeTrapAlarmResponse() (response *DescribeTrapAlarmResponse) {
	response = &DescribeTrapAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询SNMP TRAP明细告警信息
func (c *Client) DescribeTrapAlarm(request *DescribeTrapAlarmRequest) (response *DescribeTrapAlarmResponse, err error) {
	if request == nil {
		request = NewDescribeTrapAlarmRequest()
	}
	response = NewDescribeTrapAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopoSettingRequest() (request *DescribeTopoSettingRequest) {
	request = &DescribeTopoSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeTopoSetting")
	return
}

func NewDescribeTopoSettingResponse() (response *DescribeTopoSettingResponse) {
	response = &DescribeTopoSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户网络拓扑位置配置
func (c *Client) DescribeTopoSetting(request *DescribeTopoSettingRequest) (response *DescribeTopoSettingResponse, err error) {
	if request == nil {
		request = NewDescribeTopoSettingRequest()
	}
	response = NewDescribeTopoSettingResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRacksRequest() (request *ModifyRacksRequest) {
	request = &ModifyRacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyRacks")
	return
}

func NewModifyRacksResponse() (response *ModifyRacksResponse) {
	response = &ModifyRacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机架信息
func (c *Client) ModifyRacks(request *ModifyRacksRequest) (response *ModifyRacksResponse, err error) {
	if request == nil {
		request = NewModifyRacksRequest()
	}
	response = NewModifyRacksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplateConfigRequest() (request *DescribeTemplateConfigRequest) {
	request = &DescribeTemplateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeTemplateConfig")
	return
}

func NewDescribeTemplateConfigResponse() (response *DescribeTemplateConfigResponse) {
	response = &DescribeTemplateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络设备配置模板绑定
func (c *Client) DescribeTemplateConfig(request *DescribeTemplateConfigRequest) (response *DescribeTemplateConfigResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateConfigRequest()
	}
	response = NewDescribeTemplateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmStatisticsRequest() (request *DescribeAlarmStatisticsRequest) {
	request = &DescribeAlarmStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeAlarmStatistics")
	return
}

func NewDescribeAlarmStatisticsResponse() (response *DescribeAlarmStatisticsResponse) {
	response = &DescribeAlarmStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询未处理告警总数
func (c *Client) DescribeAlarmStatistics(request *DescribeAlarmStatisticsRequest) (response *DescribeAlarmStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmStatisticsRequest()
	}
	response = NewDescribeAlarmStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewGetLanDetailDataRequest() (request *GetLanDetailDataRequest) {
	request = &GetLanDetailDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetLanDetailData")
	return
}

func NewGetLanDetailDataResponse() (response *GetLanDetailDataResponse) {
	response = &GetLanDetailDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内网探测明细数据
func (c *Client) GetLanDetailData(request *GetLanDetailDataRequest) (response *GetLanDetailDataResponse, err error) {
	if request == nil {
		request = NewGetLanDetailDataRequest()
	}
	response = NewGetLanDetailDataResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNetworkDevicesRequest() (request *DeleteNetworkDevicesRequest) {
	request = &DeleteNetworkDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteNetworkDevices")
	return
}

func NewDeleteNetworkDevicesResponse() (response *DeleteNetworkDevicesResponse) {
	response = &DeleteNetworkDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络设备信息
func (c *Client) DeleteNetworkDevices(request *DeleteNetworkDevicesRequest) (response *DeleteNetworkDevicesResponse, err error) {
	if request == nil {
		request = NewDeleteNetworkDevicesRequest()
	}
	response = NewDeleteNetworkDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewGetGwConfigRequest() (request *GetGwConfigRequest) {
	request = &GetGwConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetGwConfig")
	return
}

func NewGetGwConfigResponse() (response *GetGwConfigResponse) {
	response = &GetGwConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关探测数据
func (c *Client) GetGwConfig(request *GetGwConfigRequest) (response *GetGwConfigResponse, err error) {
	if request == nil {
		request = NewGetGwConfigRequest()
	}
	response = NewGetGwConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetWanDetailDataRequest() (request *GetWanDetailDataRequest) {
	request = &GetWanDetailDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetWanDetailData")
	return
}

func NewGetWanDetailDataResponse() (response *GetWanDetailDataResponse) {
	response = &GetWanDetailDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询外网探测明细数据
func (c *Client) GetWanDetailData(request *GetWanDetailDataRequest) (response *GetWanDetailDataResponse, err error) {
	if request == nil {
		request = NewGetWanDetailDataRequest()
	}
	response = NewGetWanDetailDataResponse()
	err = c.Send(request, response)
	return
}

func NewGetLanDataRequest() (request *GetLanDataRequest) {
	request = &GetLanDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetLanData")
	return
}

func NewGetLanDataResponse() (response *GetLanDataResponse) {
	response = &GetLanDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询内网探测数据
func (c *Client) GetLanData(request *GetLanDataRequest) (response *GetLanDataResponse, err error) {
	if request == nil {
		request = NewGetLanDataRequest()
	}
	response = NewGetLanDataResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteTemplate")
	return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络设备配置模板
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteTemplateRequest()
	}
	response = NewDeleteTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetGwDetailDataRequest() (request *GetGwDetailDataRequest) {
	request = &GetGwDetailDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "GetGwDetailData")
	return
}

func NewGetGwDetailDataResponse() (response *GetGwDetailDataResponse) {
	response = &GetGwDetailDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关探测明细数据
func (c *Client) GetGwDetailData(request *GetGwDetailDataRequest) (response *GetGwDetailDataResponse, err error) {
	if request == nil {
		request = NewGetGwDetailDataRequest()
	}
	response = NewGetGwDetailDataResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIdcExportLineRequest() (request *ModifyIdcExportLineRequest) {
	request = &ModifyIdcExportLineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "ModifyIdcExportLine")
	return
}

func NewModifyIdcExportLineResponse() (response *ModifyIdcExportLineResponse) {
	response = &ModifyIdcExportLineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改IDC专线出口
func (c *Client) ModifyIdcExportLine(request *ModifyIdcExportLineRequest) (response *ModifyIdcExportLineResponse, err error) {
	if request == nil {
		request = NewModifyIdcExportLineRequest()
	}
	response = NewModifyIdcExportLineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIdcExportLineTopoRequest() (request *DescribeIdcExportLineTopoRequest) {
	request = &DescribeIdcExportLineTopoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIdcExportLineTopo")
	return
}

func NewDescribeIdcExportLineTopoResponse() (response *DescribeIdcExportLineTopoResponse) {
	response = &DescribeIdcExportLineTopoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询AZ内专线出口的拓扑连接关系
func (c *Client) DescribeIdcExportLineTopo(request *DescribeIdcExportLineTopoRequest) (response *DescribeIdcExportLineTopoResponse, err error) {
	if request == nil {
		request = NewDescribeIdcExportLineTopoRequest()
	}
	response = NewDescribeIdcExportLineTopoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewayClustersRequest() (request *DeleteGatewayClustersRequest) {
	request = &DeleteGatewayClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DeleteGatewayClusters")
	return
}

func NewDeleteGatewayClustersResponse() (response *DeleteGatewayClustersResponse) {
	response = &DeleteGatewayClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网关集群
func (c *Client) DeleteGatewayClusters(request *DeleteGatewayClustersRequest) (response *DeleteGatewayClustersResponse, err error) {
	if request == nil {
		request = NewDeleteGatewayClustersRequest()
	}
	response = NewDeleteGatewayClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIdcExportLineRequest() (request *DescribeIdcExportLineRequest) {
	request = &DescribeIdcExportLineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "DescribeIdcExportLine")
	return
}

func NewDescribeIdcExportLineResponse() (response *DescribeIdcExportLineResponse) {
	response = &DescribeIdcExportLineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询专线出口
func (c *Client) DescribeIdcExportLine(request *DescribeIdcExportLineRequest) (response *DescribeIdcExportLineResponse, err error) {
	if request == nil {
		request = NewDescribeIdcExportLineRequest()
	}
	response = NewDescribeIdcExportLineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIdcExportLineRequest() (request *CreateIdcExportLineRequest) {
	request = &CreateIdcExportLineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("nms", APIVersion, "CreateIdcExportLine")
	return
}

func NewCreateIdcExportLineResponse() (response *CreateIdcExportLineResponse) {
	response = &CreateIdcExportLineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加IDC专线出口
func (c *Client) CreateIdcExportLine(request *CreateIdcExportLineRequest) (response *CreateIdcExportLineResponse, err error) {
	if request == nil {
		request = NewCreateIdcExportLineRequest()
	}
	response = NewCreateIdcExportLineResponse()
	err = c.Send(request, response)
	return
}
