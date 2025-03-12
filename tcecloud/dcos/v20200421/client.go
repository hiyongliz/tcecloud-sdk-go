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

package v20200421

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-04-21"

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

func NewDescribeFieldsEnumExRequest() (request *DescribeFieldsEnumExRequest) {
	request = &DescribeFieldsEnumExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeFieldsEnumEx")
	return
}

func NewDescribeFieldsEnumExResponse() (response *DescribeFieldsEnumExResponse) {
	response = &DescribeFieldsEnumExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取cmdb各字段枚举值
func (c *Client) DescribeFieldsEnumEx(request *DescribeFieldsEnumExRequest) (response *DescribeFieldsEnumExResponse, err error) {
	if request == nil {
		request = NewDescribeFieldsEnumExRequest()
	}
	response = NewDescribeFieldsEnumExResponse()
	err = c.Send(request, response)
	return
}

func NewQueryTaskExRequest() (request *QueryTaskExRequest) {
	request = &QueryTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "QueryTaskEx")
	return
}

func NewQueryTaskExResponse() (response *QueryTaskExResponse) {
	response = &QueryTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询dcos装机结果
func (c *Client) QueryTaskEx(request *QueryTaskExRequest) (response *QueryTaskExResponse, err error) {
	if request == nil {
		request = NewQueryTaskExRequest()
	}
	response = NewQueryTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerQueryRequest() (request *RelocateServerQueryRequest) {
	request = &RelocateServerQueryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RelocateServerQuery")
	return
}

func NewRelocateServerQueryResponse() (response *RelocateServerQueryResponse) {
	response = &RelocateServerQueryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器搬迁状态查询接口
func (c *Client) RelocateServerQuery(request *RelocateServerQueryRequest) (response *RelocateServerQueryResponse, err error) {
	if request == nil {
		request = NewRelocateServerQueryRequest()
	}
	response = NewRelocateServerQueryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOutbandInfoExRequest() (request *DescribeOutbandInfoExRequest) {
	request = &DescribeOutbandInfoExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeOutbandInfoEx")
	return
}

func NewDescribeOutbandInfoExResponse() (response *DescribeOutbandInfoExResponse) {
	response = &DescribeOutbandInfoExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 带外管理,账号密码查询
func (c *Client) DescribeOutbandInfoEx(request *DescribeOutbandInfoExRequest) (response *DescribeOutbandInfoExResponse, err error) {
	if request == nil {
		request = NewDescribeOutbandInfoExRequest()
	}
	response = NewDescribeOutbandInfoExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerRelatedNetdevicesRequest() (request *DescribeServerRelatedNetdevicesRequest) {
	request = &DescribeServerRelatedNetdevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeServerRelatedNetdevices")
	return
}

func NewDescribeServerRelatedNetdevicesResponse() (response *DescribeServerRelatedNetdevicesResponse) {
	response = &DescribeServerRelatedNetdevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务器关联网络设备
func (c *Client) DescribeServerRelatedNetdevices(request *DescribeServerRelatedNetdevicesRequest) (response *DescribeServerRelatedNetdevicesResponse, err error) {
	if request == nil {
		request = NewDescribeServerRelatedNetdevicesRequest()
	}
	response = NewDescribeServerRelatedNetdevicesResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateVMWanIPListRequest() (request *AllocateVMWanIPListRequest) {
	request = &AllocateVMWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AllocateVMWanIPList")
	return
}

func NewAllocateVMWanIPListResponse() (response *AllocateVMWanIPListResponse) {
	response = &AllocateVMWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配VM外网IP
func (c *Client) AllocateVMWanIPList(request *AllocateVMWanIPListRequest) (response *AllocateVMWanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateVMWanIPListRequest()
	}
	response = NewAllocateVMWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewServerSpecialAllocateLanIPExRequest() (request *ServerSpecialAllocateLanIPExRequest) {
	request = &ServerSpecialAllocateLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ServerSpecialAllocateLanIPEx")
	return
}

func NewServerSpecialAllocateLanIPExResponse() (response *ServerSpecialAllocateLanIPExResponse) {
	response = &ServerSpecialAllocateLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器批量分配相应虚拟内网段IP接口
func (c *Client) ServerSpecialAllocateLanIPEx(request *ServerSpecialAllocateLanIPExRequest) (response *ServerSpecialAllocateLanIPExResponse, err error) {
	if request == nil {
		request = NewServerSpecialAllocateLanIPExRequest()
	}
	response = NewServerSpecialAllocateLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerFinishRequest() (request *RelocateServerFinishRequest) {
	request = &RelocateServerFinishRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RelocateServerFinish")
	return
}

func NewRelocateServerFinishResponse() (response *RelocateServerFinishResponse) {
	response = &RelocateServerFinishResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器搬迁完成接口
func (c *Client) RelocateServerFinish(request *RelocateServerFinishRequest) (response *RelocateServerFinishResponse, err error) {
	if request == nil {
		request = NewRelocateServerFinishRequest()
	}
	response = NewRelocateServerFinishResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNetdeviceRelatedServersExRequest() (request *DescribeNetdeviceRelatedServersExRequest) {
	request = &DescribeNetdeviceRelatedServersExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeNetdeviceRelatedServersEx")
	return
}

func NewDescribeNetdeviceRelatedServersExResponse() (response *DescribeNetdeviceRelatedServersExResponse) {
	response = &DescribeNetdeviceRelatedServersExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取网络设备关联的服务器
func (c *Client) DescribeNetdeviceRelatedServersEx(request *DescribeNetdeviceRelatedServersExRequest) (response *DescribeNetdeviceRelatedServersExResponse, err error) {
	if request == nil {
		request = NewDescribeNetdeviceRelatedServersExRequest()
	}
	response = NewDescribeNetdeviceRelatedServersExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerHarddisksRequest() (request *DescribeServerHarddisksRequest) {
	request = &DescribeServerHarddisksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeServerHarddisks")
	return
}

func NewDescribeServerHarddisksResponse() (response *DescribeServerHarddisksResponse) {
	response = &DescribeServerHarddisksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取物理机磁盘列表
func (c *Client) DescribeServerHarddisks(request *DescribeServerHarddisksRequest) (response *DescribeServerHarddisksResponse, err error) {
	if request == nil {
		request = NewDescribeServerHarddisksRequest()
	}
	response = NewDescribeServerHarddisksResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRAIDRequest() (request *DeleteRAIDRequest) {
	request = &DeleteRAIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteRAID")
	return
}

func NewDeleteRAIDResponse() (response *DeleteRAIDResponse) {
	response = &DeleteRAIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RAID
func (c *Client) DeleteRAID(request *DeleteRAIDRequest) (response *DeleteRAIDResponse, err error) {
	if request == nil {
		request = NewDeleteRAIDRequest()
	}
	response = NewDeleteRAIDResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLabelRequest() (request *DeleteLabelRequest) {
	request = &DeleteLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteLabel")
	return
}

func NewDeleteLabelResponse() (response *DeleteLabelResponse) {
	response = &DeleteLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--删除标签
func (c *Client) DeleteLabel(request *DeleteLabelRequest) (response *DeleteLabelResponse, err error) {
	if request == nil {
		request = NewDeleteLabelRequest()
	}
	response = NewDeleteLabelResponse()
	err = c.Send(request, response)
	return
}

func NewAddRAIDRequest() (request *AddRAIDRequest) {
	request = &AddRAIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AddRAID")
	return
}

func NewAddRAIDResponse() (response *AddRAIDResponse) {
	response = &AddRAIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加RAID
func (c *Client) AddRAID(request *AddRAIDRequest) (response *AddRAIDResponse, err error) {
	if request == nil {
		request = NewAddRAIDRequest()
	}
	response = NewAddRAIDResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomScriptSetsRequest() (request *DescribeCustomScriptSetsRequest) {
	request = &DescribeCustomScriptSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeCustomScriptSets")
	return
}

func NewDescribeCustomScriptSetsResponse() (response *DescribeCustomScriptSetsResponse) {
	response = &DescribeCustomScriptSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自定义脚本集信息
func (c *Client) DescribeCustomScriptSets(request *DescribeCustomScriptSetsRequest) (response *DescribeCustomScriptSetsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomScriptSetsRequest()
	}
	response = NewDescribeCustomScriptSetsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomScriptsRequest() (request *DeleteCustomScriptsRequest) {
	request = &DeleteCustomScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteCustomScripts")
	return
}

func NewDeleteCustomScriptsResponse() (response *DeleteCustomScriptsResponse) {
	response = &DeleteCustomScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户自定义脚本
func (c *Client) DeleteCustomScripts(request *DeleteCustomScriptsRequest) (response *DeleteCustomScriptsResponse, err error) {
	if request == nil {
		request = NewDeleteCustomScriptsRequest()
	}
	response = NewDeleteCustomScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOSListRequest() (request *DescribeOSListRequest) {
	request = &DescribeOSListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeOSList")
	return
}

func NewDescribeOSListResponse() (response *DescribeOSListResponse) {
	response = &DescribeOSListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeOSList
func (c *Client) DescribeOSList(request *DescribeOSListRequest) (response *DescribeOSListResponse, err error) {
	if request == nil {
		request = NewDescribeOSListRequest()
	}
	response = NewDescribeOSListResponse()
	err = c.Send(request, response)
	return
}

func NewRenameServersRequest() (request *RenameServersRequest) {
	request = &RenameServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RenameServers")
	return
}

func NewRenameServersResponse() (response *RenameServersResponse) {
	response = &RenameServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重命名服务器
func (c *Client) RenameServers(request *RenameServersRequest) (response *RenameServersResponse, err error) {
	if request == nil {
		request = NewRenameServersRequest()
	}
	response = NewRenameServersResponse()
	err = c.Send(request, response)
	return
}

func NewServerRecycleLanIPExRequest() (request *ServerRecycleLanIPExRequest) {
	request = &ServerRecycleLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ServerRecycleLanIPEx")
	return
}

func NewServerRecycleLanIPExResponse() (response *ServerRecycleLanIPExResponse) {
	response = &ServerRecycleLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器回收内网IP接口
func (c *Client) ServerRecycleLanIPEx(request *ServerRecycleLanIPExRequest) (response *ServerRecycleLanIPExResponse, err error) {
	if request == nil {
		request = NewServerRecycleLanIPExRequest()
	}
	response = NewServerRecycleLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewWithdrawServerExRequest() (request *WithdrawServerExRequest) {
	request = &WithdrawServerExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "WithdrawServerEx")
	return
}

func NewWithdrawServerExResponse() (response *WithdrawServerExResponse) {
	response = &WithdrawServerExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下架物理服务器
func (c *Client) WithdrawServerEx(request *WithdrawServerExRequest) (response *WithdrawServerExResponse, err error) {
	if request == nil {
		request = NewWithdrawServerExRequest()
	}
	response = NewWithdrawServerExResponse()
	err = c.Send(request, response)
	return
}

func NewModifyImageRequest() (request *ModifyImageRequest) {
	request = &ModifyImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyImage")
	return
}

func NewModifyImageResponse() (response *ModifyImageResponse) {
	response = &ModifyImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改镜像
func (c *Client) ModifyImage(request *ModifyImageRequest) (response *ModifyImageResponse, err error) {
	if request == nil {
		request = NewModifyImageRequest()
	}
	response = NewModifyImageResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleVMLanIPListRequest() (request *RecycleVMLanIPListRequest) {
	request = &RecycleVMLanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RecycleVMLanIPList")
	return
}

func NewRecycleVMLanIPListResponse() (response *RecycleVMLanIPListResponse) {
	response = &RecycleVMLanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收VM内网IP
func (c *Client) RecycleVMLanIPList(request *RecycleVMLanIPListRequest) (response *RecycleVMLanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleVMLanIPListRequest()
	}
	response = NewRecycleVMLanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewConfigServerDefRequest() (request *ConfigServerDefRequest) {
	request = &ConfigServerDefRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ConfigServerDef")
	return
}

func NewConfigServerDefResponse() (response *ConfigServerDefResponse) {
	response = &ConfigServerDefResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置初始化
func (c *Client) ConfigServerDef(request *ConfigServerDefRequest) (response *ConfigServerDefResponse, err error) {
	if request == nil {
		request = NewConfigServerDefRequest()
	}
	response = NewConfigServerDefResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRAIDRequest() (request *ModifyRAIDRequest) {
	request = &ModifyRAIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyRAID")
	return
}

func NewModifyRAIDResponse() (response *ModifyRAIDResponse) {
	response = &ModifyRAIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改RAID
func (c *Client) ModifyRAID(request *ModifyRAIDRequest) (response *ModifyRAIDResponse, err error) {
	if request == nil {
		request = NewModifyRAIDRequest()
	}
	response = NewModifyRAIDResponse()
	err = c.Send(request, response)
	return
}

func NewServerSpecialRecycleLanIPExRequest() (request *ServerSpecialRecycleLanIPExRequest) {
	request = &ServerSpecialRecycleLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ServerSpecialRecycleLanIPEx")
	return
}

func NewServerSpecialRecycleLanIPExResponse() (response *ServerSpecialRecycleLanIPExResponse) {
	response = &ServerSpecialRecycleLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器批量回收相应虚拟内网段IP接口
func (c *Client) ServerSpecialRecycleLanIPEx(request *ServerSpecialRecycleLanIPExRequest) (response *ServerSpecialRecycleLanIPExResponse, err error) {
	if request == nil {
		request = NewServerSpecialRecycleLanIPExRequest()
	}
	response = NewServerSpecialRecycleLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewAddLabelRequest() (request *AddLabelRequest) {
	request = &AddLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AddLabel")
	return
}

func NewAddLabelResponse() (response *AddLabelResponse) {
	response = &AddLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--新增标签
func (c *Client) AddLabel(request *AddLabelRequest) (response *AddLabelResponse, err error) {
	if request == nil {
		request = NewAddLabelRequest()
	}
	response = NewAddLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePhysvrsOverviewRequest() (request *DescribePhysvrsOverviewRequest) {
	request = &DescribePhysvrsOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribePhysvrsOverview")
	return
}

func NewDescribePhysvrsOverviewResponse() (response *DescribePhysvrsOverviewResponse) {
	response = &DescribePhysvrsOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器资源概览
func (c *Client) DescribePhysvrsOverview(request *DescribePhysvrsOverviewRequest) (response *DescribePhysvrsOverviewResponse, err error) {
	if request == nil {
		request = NewDescribePhysvrsOverviewRequest()
	}
	response = NewDescribePhysvrsOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOutbandStrategyRequest() (request *DeleteOutbandStrategyRequest) {
	request = &DeleteOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteOutbandStrategy")
	return
}

func NewDeleteOutbandStrategyResponse() (response *DeleteOutbandStrategyResponse) {
	response = &DeleteOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义带外密码策略
func (c *Client) DeleteOutbandStrategy(request *DeleteOutbandStrategyRequest) (response *DeleteOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewDeleteOutbandStrategyRequest()
	}
	response = NewDeleteOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerStartRequest() (request *RelocateServerStartRequest) {
	request = &RelocateServerStartRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RelocateServerStart")
	return
}

func NewRelocateServerStartResponse() (response *RelocateServerStartResponse) {
	response = &RelocateServerStartResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器搬迁发起接口
func (c *Client) RelocateServerStart(request *RelocateServerStartRequest) (response *RelocateServerStartResponse, err error) {
	if request == nil {
		request = NewRelocateServerStartRequest()
	}
	response = NewRelocateServerStartResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLabelRequest() (request *DescribeLabelRequest) {
	request = &DescribeLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeLabel")
	return
}

func NewDescribeLabelResponse() (response *DescribeLabelResponse) {
	response = &DescribeLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--查询标签
func (c *Client) DescribeLabel(request *DescribeLabelRequest) (response *DescribeLabelResponse, err error) {
	if request == nil {
		request = NewDescribeLabelRequest()
	}
	response = NewDescribeLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePhysvrsListExRequest() (request *DescribePhysvrsListExRequest) {
	request = &DescribePhysvrsListExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribePhysvrsListEx")
	return
}

func NewDescribePhysvrsListExResponse() (response *DescribePhysvrsListExResponse) {
	response = &DescribePhysvrsListExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取物理服务器列表
func (c *Client) DescribePhysvrsListEx(request *DescribePhysvrsListExRequest) (response *DescribePhysvrsListExResponse, err error) {
	if request == nil {
		request = NewDescribePhysvrsListExRequest()
	}
	response = NewDescribePhysvrsListExResponse()
	err = c.Send(request, response)
	return
}

func NewAddOutbandStrategyRequest() (request *AddOutbandStrategyRequest) {
	request = &AddOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AddOutbandStrategy")
	return
}

func NewAddOutbandStrategyResponse() (response *AddOutbandStrategyResponse) {
	response = &AddOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加自定义带外密码策略
func (c *Client) AddOutbandStrategy(request *AddOutbandStrategyRequest) (response *AddOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewAddOutbandStrategyRequest()
	}
	response = NewAddOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCustomScriptSetsRequest() (request *DeleteCustomScriptSetsRequest) {
	request = &DeleteCustomScriptSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteCustomScriptSets")
	return
}

func NewDeleteCustomScriptSetsResponse() (response *DeleteCustomScriptSetsResponse) {
	response = &DeleteCustomScriptSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自定义脚本集
func (c *Client) DeleteCustomScriptSets(request *DeleteCustomScriptSetsRequest) (response *DeleteCustomScriptSetsResponse, err error) {
	if request == nil {
		request = NewDeleteCustomScriptSetsRequest()
	}
	response = NewDeleteCustomScriptSetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIgniterImagesRequest() (request *DescribeIgniterImagesRequest) {
	request = &DescribeIgniterImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeIgniterImages")
	return
}

func NewDescribeIgniterImagesResponse() (response *DescribeIgniterImagesResponse) {
	response = &DescribeIgniterImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Igniter装机镜像列表
func (c *Client) DescribeIgniterImages(request *DescribeIgniterImagesRequest) (response *DescribeIgniterImagesResponse, err error) {
	if request == nil {
		request = NewDescribeIgniterImagesRequest()
	}
	response = NewDescribeIgniterImagesResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleServerVirtualIPRequest() (request *RecycleServerVirtualIPRequest) {
	request = &RecycleServerVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RecycleServerVirtualIP")
	return
}

func NewRecycleServerVirtualIPResponse() (response *RecycleServerVirtualIPResponse) {
	response = &RecycleServerVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收物理服务器虚拟内外网ip
func (c *Client) RecycleServerVirtualIP(request *RecycleServerVirtualIPRequest) (response *RecycleServerVirtualIPResponse, err error) {
	if request == nil {
		request = NewRecycleServerVirtualIPRequest()
	}
	response = NewRecycleServerVirtualIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllVlanIdsExRequest() (request *DescribeAllVlanIdsExRequest) {
	request = &DescribeAllVlanIdsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeAllVlanIdsEx")
	return
}

func NewDescribeAllVlanIdsExResponse() (response *DescribeAllVlanIdsExResponse) {
	response = &DescribeAllVlanIdsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取所有vlanid接口
func (c *Client) DescribeAllVlanIdsEx(request *DescribeAllVlanIdsExRequest) (response *DescribeAllVlanIdsExResponse, err error) {
	if request == nil {
		request = NewDescribeAllVlanIdsExRequest()
	}
	response = NewDescribeAllVlanIdsExResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleVMVirtualIPRequest() (request *RecycleVMVirtualIPRequest) {
	request = &RecycleVMVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RecycleVMVirtualIP")
	return
}

func NewRecycleVMVirtualIPResponse() (response *RecycleVMVirtualIPResponse) {
	response = &RecycleVMVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收VM虚拟内外网IP
func (c *Client) RecycleVMVirtualIP(request *RecycleVMVirtualIPRequest) (response *RecycleVMVirtualIPResponse, err error) {
	if request == nil {
		request = NewRecycleVMVirtualIPRequest()
	}
	response = NewRecycleVMVirtualIPResponse()
	err = c.Send(request, response)
	return
}

func NewImportCmdbServersRequest() (request *ImportCmdbServersRequest) {
	request = &ImportCmdbServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ImportCmdbServers")
	return
}

func NewImportCmdbServersResponse() (response *ImportCmdbServersResponse) {
	response = &ImportCmdbServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入物理服务器
func (c *Client) ImportCmdbServers(request *ImportCmdbServersRequest) (response *ImportCmdbServersResponse, err error) {
	if request == nil {
		request = NewImportCmdbServersRequest()
	}
	response = NewImportCmdbServersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyServersRequest() (request *ModifyServersRequest) {
	request = &ModifyServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyServers")
	return
}

func NewModifyServersResponse() (response *ModifyServersResponse) {
	response = &ModifyServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新服务器信息请求信息
func (c *Client) ModifyServers(request *ModifyServersRequest) (response *ModifyServersResponse, err error) {
	if request == nil {
		request = NewModifyServersRequest()
	}
	response = NewModifyServersResponse()
	err = c.Send(request, response)
	return
}

func NewAddImageRequest() (request *AddImageRequest) {
	request = &AddImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AddImage")
	return
}

func NewAddImageResponse() (response *AddImageResponse) {
	response = &AddImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增镜像
func (c *Client) AddImage(request *AddImageRequest) (response *AddImageResponse, err error) {
	if request == nil {
		request = NewAddImageRequest()
	}
	response = NewAddImageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomScriptRequest() (request *ModifyCustomScriptRequest) {
	request = &ModifyCustomScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyCustomScript")
	return
}

func NewModifyCustomScriptResponse() (response *ModifyCustomScriptResponse) {
	response = &ModifyCustomScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自定义脚本信息
func (c *Client) ModifyCustomScript(request *ModifyCustomScriptRequest) (response *ModifyCustomScriptResponse, err error) {
	if request == nil {
		request = NewModifyCustomScriptRequest()
	}
	response = NewModifyCustomScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRelocateInfoRequest() (request *DescribeRelocateInfoRequest) {
	request = &DescribeRelocateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeRelocateInfo")
	return
}

func NewDescribeRelocateInfoResponse() (response *DescribeRelocateInfoResponse) {
	response = &DescribeRelocateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询搬迁目的机架机位信息
func (c *Client) DescribeRelocateInfo(request *DescribeRelocateInfoRequest) (response *DescribeRelocateInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRelocateInfoRequest()
	}
	response = NewDescribeRelocateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteVMListRequest() (request *DeleteVMListRequest) {
	request = &DeleteVMListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteVMList")
	return
}

func NewDeleteVMListResponse() (response *DeleteVMListResponse) {
	response = &DeleteVMListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除虚拟机
func (c *Client) DeleteVMList(request *DeleteVMListRequest) (response *DeleteVMListResponse, err error) {
	if request == nil {
		request = NewDeleteVMListRequest()
	}
	response = NewDeleteVMListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServersRequest() (request *DeleteServersRequest) {
	request = &DeleteServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteServers")
	return
}

func NewDeleteServersResponse() (response *DeleteServersResponse) {
	response = &DeleteServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除服务器信息
func (c *Client) DeleteServers(request *DeleteServersRequest) (response *DeleteServersResponse, err error) {
	if request == nil {
		request = NewDeleteServersRequest()
	}
	response = NewDeleteServersResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBackendComponentsExRequest() (request *QueryBackendComponentsExRequest) {
	request = &QueryBackendComponentsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "QueryBackendComponentsEx")
	return
}

func NewQueryBackendComponentsExResponse() (response *QueryBackendComponentsExResponse) {
	response = &QueryBackendComponentsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询后端组件接口
func (c *Client) QueryBackendComponentsEx(request *QueryBackendComponentsExRequest) (response *QueryBackendComponentsExResponse, err error) {
	if request == nil {
		request = NewQueryBackendComponentsExRequest()
	}
	response = NewQueryBackendComponentsExResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleServerWanIPListRequest() (request *RecycleServerWanIPListRequest) {
	request = &RecycleServerWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RecycleServerWanIPList")
	return
}

func NewRecycleServerWanIPListResponse() (response *RecycleServerWanIPListResponse) {
	response = &RecycleServerWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收服务器外网IP
func (c *Client) RecycleServerWanIPList(request *RecycleServerWanIPListRequest) (response *RecycleServerWanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleServerWanIPListRequest()
	}
	response = NewRecycleServerWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomScriptsRequest() (request *DescribeCustomScriptsRequest) {
	request = &DescribeCustomScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeCustomScripts")
	return
}

func NewDescribeCustomScriptsResponse() (response *DescribeCustomScriptsResponse) {
	response = &DescribeCustomScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户自定义脚本
func (c *Client) DescribeCustomScripts(request *DescribeCustomScriptsRequest) (response *DescribeCustomScriptsResponse, err error) {
	if request == nil {
		request = NewDescribeCustomScriptsRequest()
	}
	response = NewDescribeCustomScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateVMLanIPListRequest() (request *AllocateVMLanIPListRequest) {
	request = &AllocateVMLanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AllocateVMLanIPList")
	return
}

func NewAllocateVMLanIPListResponse() (response *AllocateVMLanIPListResponse) {
	response = &AllocateVMLanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配VM内网IP
func (c *Client) AllocateVMLanIPList(request *AllocateVMLanIPListRequest) (response *AllocateVMLanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateVMLanIPListRequest()
	}
	response = NewAllocateVMLanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLabelsRequest() (request *ModifyLabelsRequest) {
	request = &ModifyLabelsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyLabels")
	return
}

func NewModifyLabelsResponse() (response *ModifyLabelsResponse) {
	response = &ModifyLabelsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改物理服务器标签
func (c *Client) ModifyLabels(request *ModifyLabelsRequest) (response *ModifyLabelsResponse, err error) {
	if request == nil {
		request = NewModifyLabelsRequest()
	}
	response = NewModifyLabelsResponse()
	err = c.Send(request, response)
	return
}

func NewQueryOutbandStrategyRequest() (request *QueryOutbandStrategyRequest) {
	request = &QueryOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "QueryOutbandStrategy")
	return
}

func NewQueryOutbandStrategyResponse() (response *QueryOutbandStrategyResponse) {
	response = &QueryOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询自定义带外密码策略
func (c *Client) QueryOutbandStrategy(request *QueryOutbandStrategyRequest) (response *QueryOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewQueryOutbandStrategyRequest()
	}
	response = NewQueryOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleServerSpecialWanIPListRequest() (request *RecycleServerSpecialWanIPListRequest) {
	request = &RecycleServerSpecialWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RecycleServerSpecialWanIPList")
	return
}

func NewRecycleServerSpecialWanIPListResponse() (response *RecycleServerSpecialWanIPListResponse) {
	response = &RecycleServerSpecialWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收服务器特殊外网IP
func (c *Client) RecycleServerSpecialWanIPList(request *RecycleServerSpecialWanIPListRequest) (response *RecycleServerSpecialWanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleServerSpecialWanIPListRequest()
	}
	response = NewRecycleServerSpecialWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImagePathRequest() (request *DescribeImagePathRequest) {
	request = &DescribeImagePathRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeImagePath")
	return
}

func NewDescribeImagePathResponse() (response *DescribeImagePathResponse) {
	response = &DescribeImagePathResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像原始地址接口
func (c *Client) DescribeImagePath(request *DescribeImagePathRequest) (response *DescribeImagePathResponse, err error) {
	if request == nil {
		request = NewDescribeImagePathRequest()
	}
	response = NewDescribeImagePathResponse()
	err = c.Send(request, response)
	return
}

func NewReinstallOsExRequest() (request *ReinstallOsExRequest) {
	request = &ReinstallOsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ReinstallOsEx")
	return
}

func NewReinstallOsExResponse() (response *ReinstallOsExResponse) {
	response = &ReinstallOsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// OS部署 部署系统
func (c *Client) ReinstallOsEx(request *ReinstallOsExRequest) (response *ReinstallOsExResponse, err error) {
	if request == nil {
		request = NewReinstallOsExRequest()
	}
	response = NewReinstallOsExResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServersRequest() (request *DescribeServersRequest) {
	request = &DescribeServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeServers")
	return
}

func NewDescribeServersResponse() (response *DescribeServersResponse) {
	response = &DescribeServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务器信息
func (c *Client) DescribeServers(request *DescribeServersRequest) (response *DescribeServersResponse, err error) {
	if request == nil {
		request = NewDescribeServersRequest()
	}
	response = NewDescribeServersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOsDeployDebugInfoExRequest() (request *DescribeOsDeployDebugInfoExRequest) {
	request = &DescribeOsDeployDebugInfoExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeOsDeployDebugInfoEx")
	return
}

func NewDescribeOsDeployDebugInfoExResponse() (response *DescribeOsDeployDebugInfoExResponse) {
	response = &DescribeOsDeployDebugInfoExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询os部署调试信息
func (c *Client) DescribeOsDeployDebugInfoEx(request *DescribeOsDeployDebugInfoExRequest) (response *DescribeOsDeployDebugInfoExResponse, err error) {
	if request == nil {
		request = NewDescribeOsDeployDebugInfoExRequest()
	}
	response = NewDescribeOsDeployDebugInfoExResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageRequest() (request *DeleteImageRequest) {
	request = &DeleteImageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DeleteImage")
	return
}

func NewDeleteImageResponse() (response *DeleteImageResponse) {
	response = &DeleteImageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像
func (c *Client) DeleteImage(request *DeleteImageRequest) (response *DeleteImageResponse, err error) {
	if request == nil {
		request = NewDeleteImageRequest()
	}
	response = NewDeleteImageResponse()
	err = c.Send(request, response)
	return
}

func NewServerAllocateLanIPExRequest() (request *ServerAllocateLanIPExRequest) {
	request = &ServerAllocateLanIPExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ServerAllocateLanIPEx")
	return
}

func NewServerAllocateLanIPExResponse() (response *ServerAllocateLanIPExResponse) {
	response = &ServerAllocateLanIPExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器分配内网IP
func (c *Client) ServerAllocateLanIPEx(request *ServerAllocateLanIPExRequest) (response *ServerAllocateLanIPExResponse, err error) {
	if request == nil {
		request = NewServerAllocateLanIPExRequest()
	}
	response = NewServerAllocateLanIPExResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomScriptRequest() (request *CreateCustomScriptRequest) {
	request = &CreateCustomScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "CreateCustomScript")
	return
}

func NewCreateCustomScriptResponse() (response *CreateCustomScriptResponse) {
	response = &CreateCustomScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用户自定义脚本，用于DCOS装机后置脚本的执行。
func (c *Client) CreateCustomScript(request *CreateCustomScriptRequest) (response *CreateCustomScriptResponse, err error) {
	if request == nil {
		request = NewCreateCustomScriptRequest()
	}
	response = NewCreateCustomScriptResponse()
	err = c.Send(request, response)
	return
}

func NewGetImageFieldsEnumRequest() (request *GetImageFieldsEnumRequest) {
	request = &GetImageFieldsEnumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "GetImageFieldsEnum")
	return
}

func NewGetImageFieldsEnumResponse() (response *GetImageFieldsEnumResponse) {
	response = &GetImageFieldsEnumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取镜像管理字段枚举值
func (c *Client) GetImageFieldsEnum(request *GetImageFieldsEnumRequest) (response *GetImageFieldsEnumResponse, err error) {
	if request == nil {
		request = NewGetImageFieldsEnumRequest()
	}
	response = NewGetImageFieldsEnumResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRelocateServerHistoryRequest() (request *DescribeRelocateServerHistoryRequest) {
	request = &DescribeRelocateServerHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeRelocateServerHistory")
	return
}

func NewDescribeRelocateServerHistoryResponse() (response *DescribeRelocateServerHistoryResponse) {
	response = &DescribeRelocateServerHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询历史搬迁接口
func (c *Client) DescribeRelocateServerHistory(request *DescribeRelocateServerHistoryRequest) (response *DescribeRelocateServerHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeRelocateServerHistoryRequest()
	}
	response = NewDescribeRelocateServerHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRAIDListRequest() (request *DescribeRAIDListRequest) {
	request = &DescribeRAIDListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeRAIDList")
	return
}

func NewDescribeRAIDListResponse() (response *DescribeRAIDListResponse) {
	response = &DescribeRAIDListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询raid
func (c *Client) DescribeRAIDList(request *DescribeRAIDListRequest) (response *DescribeRAIDListResponse, err error) {
	if request == nil {
		request = NewDescribeRAIDListRequest()
	}
	response = NewDescribeRAIDListResponse()
	err = c.Send(request, response)
	return
}

func NewAddServersRequest() (request *AddServersRequest) {
	request = &AddServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AddServers")
	return
}

func NewAddServersResponse() (response *AddServersResponse) {
	response = &AddServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加服务器信息
func (c *Client) AddServers(request *AddServersRequest) (response *AddServersResponse, err error) {
	if request == nil {
		request = NewAddServersRequest()
	}
	response = NewAddServersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeploySubtaskStepsExRequest() (request *DescribeDeploySubtaskStepsExRequest) {
	request = &DescribeDeploySubtaskStepsExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeDeploySubtaskStepsEx")
	return
}

func NewDescribeDeploySubtaskStepsExResponse() (response *DescribeDeploySubtaskStepsExResponse) {
	response = &DescribeDeploySubtaskStepsExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看部署各节点状态
func (c *Client) DescribeDeploySubtaskStepsEx(request *DescribeDeploySubtaskStepsExRequest) (response *DescribeDeploySubtaskStepsExResponse, err error) {
	if request == nil {
		request = NewDescribeDeploySubtaskStepsExRequest()
	}
	response = NewDescribeDeploySubtaskStepsExResponse()
	err = c.Send(request, response)
	return
}

func NewAddVMListRequest() (request *AddVMListRequest) {
	request = &AddVMListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AddVMList")
	return
}

func NewAddVMListResponse() (response *AddVMListResponse) {
	response = &AddVMListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加虚拟机
func (c *Client) AddVMList(request *AddVMListRequest) (response *AddVMListResponse, err error) {
	if request == nil {
		request = NewAddVMListRequest()
	}
	response = NewAddVMListResponse()
	err = c.Send(request, response)
	return
}

func NewPerformServerTaskExRequest() (request *PerformServerTaskExRequest) {
	request = &PerformServerTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "PerformServerTaskEx")
	return
}

func NewPerformServerTaskExResponse() (response *PerformServerTaskExResponse) {
	response = &PerformServerTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 带外管理 开机 关机 重启
func (c *Client) PerformServerTaskEx(request *PerformServerTaskExRequest) (response *PerformServerTaskExResponse, err error) {
	if request == nil {
		request = NewPerformServerTaskExRequest()
	}
	response = NewPerformServerTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLabelRequest() (request *ModifyLabelRequest) {
	request = &ModifyLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyLabel")
	return
}

func NewModifyLabelResponse() (response *ModifyLabelResponse) {
	response = &ModifyLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签管理--修改标签
func (c *Client) ModifyLabel(request *ModifyLabelRequest) (response *ModifyLabelResponse, err error) {
	if request == nil {
		request = NewModifyLabelRequest()
	}
	response = NewModifyLabelResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateServerSpecialWanIPListRequest() (request *AllocateServerSpecialWanIPListRequest) {
	request = &AllocateServerSpecialWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AllocateServerSpecialWanIPList")
	return
}

func NewAllocateServerSpecialWanIPListResponse() (response *AllocateServerSpecialWanIPListResponse) {
	response = &AllocateServerSpecialWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配服务器特殊外网IP
func (c *Client) AllocateServerSpecialWanIPList(request *AllocateServerSpecialWanIPListRequest) (response *AllocateServerSpecialWanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateServerSpecialWanIPListRequest()
	}
	response = NewAllocateServerSpecialWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateServerWanIPListRequest() (request *AllocateServerWanIPListRequest) {
	request = &AllocateServerWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AllocateServerWanIPList")
	return
}

func NewAllocateServerWanIPListResponse() (response *AllocateServerWanIPListResponse) {
	response = &AllocateServerWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配服务器外网IP
func (c *Client) AllocateServerWanIPList(request *AllocateServerWanIPListRequest) (response *AllocateServerWanIPListResponse, err error) {
	if request == nil {
		request = NewAllocateServerWanIPListRequest()
	}
	response = NewAllocateServerWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCustomScriptSetRequest() (request *CreateCustomScriptSetRequest) {
	request = &CreateCustomScriptSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "CreateCustomScriptSet")
	return
}

func NewCreateCustomScriptSetResponse() (response *CreateCustomScriptSetResponse) {
	response = &CreateCustomScriptSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自定义脚本集
func (c *Client) CreateCustomScriptSet(request *CreateCustomScriptSetRequest) (response *CreateCustomScriptSetResponse, err error) {
	if request == nil {
		request = NewCreateCustomScriptSetRequest()
	}
	response = NewCreateCustomScriptSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllOSListRequest() (request *DescribeAllOSListRequest) {
	request = &DescribeAllOSListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeAllOSList")
	return
}

func NewDescribeAllOSListResponse() (response *DescribeAllOSListResponse) {
	response = &DescribeAllOSListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有镜像列表
func (c *Client) DescribeAllOSList(request *DescribeAllOSListRequest) (response *DescribeAllOSListResponse, err error) {
	if request == nil {
		request = NewDescribeAllOSListRequest()
	}
	response = NewDescribeAllOSListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageListRequest() (request *DescribeImageListRequest) {
	request = &DescribeImageListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "DescribeImageList")
	return
}

func NewDescribeImageListResponse() (response *DescribeImageListResponse) {
	response = &DescribeImageListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像
func (c *Client) DescribeImageList(request *DescribeImageListRequest) (response *DescribeImageListResponse, err error) {
	if request == nil {
		request = NewDescribeImageListRequest()
	}
	response = NewDescribeImageListResponse()
	err = c.Send(request, response)
	return
}

func NewRelocateServerCancelRequest() (request *RelocateServerCancelRequest) {
	request = &RelocateServerCancelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RelocateServerCancel")
	return
}

func NewRelocateServerCancelResponse() (response *RelocateServerCancelResponse) {
	response = &RelocateServerCancelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 物理服务器取消搬迁接口
func (c *Client) RelocateServerCancel(request *RelocateServerCancelRequest) (response *RelocateServerCancelResponse, err error) {
	if request == nil {
		request = NewRelocateServerCancelRequest()
	}
	response = NewRelocateServerCancelResponse()
	err = c.Send(request, response)
	return
}

func NewRecycleVMWanIPListRequest() (request *RecycleVMWanIPListRequest) {
	request = &RecycleVMWanIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "RecycleVMWanIPList")
	return
}

func NewRecycleVMWanIPListResponse() (response *RecycleVMWanIPListResponse) {
	response = &RecycleVMWanIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回收VM外网IP
func (c *Client) RecycleVMWanIPList(request *RecycleVMWanIPListRequest) (response *RecycleVMWanIPListResponse, err error) {
	if request == nil {
		request = NewRecycleVMWanIPListRequest()
	}
	response = NewRecycleVMWanIPListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCustomScriptSetRequest() (request *ModifyCustomScriptSetRequest) {
	request = &ModifyCustomScriptSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyCustomScriptSet")
	return
}

func NewModifyCustomScriptSetResponse() (response *ModifyCustomScriptSetResponse) {
	response = &ModifyCustomScriptSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑自定义脚本集信息
func (c *Client) ModifyCustomScriptSet(request *ModifyCustomScriptSetRequest) (response *ModifyCustomScriptSetResponse, err error) {
	if request == nil {
		request = NewModifyCustomScriptSetRequest()
	}
	response = NewModifyCustomScriptSetResponse()
	err = c.Send(request, response)
	return
}

func NewSocPerformServerTaskExRequest() (request *SocPerformServerTaskExRequest) {
	request = &SocPerformServerTaskExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "SocPerformServerTaskEx")
	return
}

func NewSocPerformServerTaskExResponse() (response *SocPerformServerTaskExResponse) {
	response = &SocPerformServerTaskExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 黑石服务器带外操作 开机 关机 重启
func (c *Client) SocPerformServerTaskEx(request *SocPerformServerTaskExRequest) (response *SocPerformServerTaskExResponse, err error) {
	if request == nil {
		request = NewSocPerformServerTaskExRequest()
	}
	response = NewSocPerformServerTaskExResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateServerVirtualIPRequest() (request *AllocateServerVirtualIPRequest) {
	request = &AllocateServerVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AllocateServerVirtualIP")
	return
}

func NewAllocateServerVirtualIPResponse() (response *AllocateServerVirtualIPResponse) {
	response = &AllocateServerVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配物理服务器虚拟内外网ip
func (c *Client) AllocateServerVirtualIP(request *AllocateServerVirtualIPRequest) (response *AllocateServerVirtualIPResponse, err error) {
	if request == nil {
		request = NewAllocateServerVirtualIPRequest()
	}
	response = NewAllocateServerVirtualIPResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOutbandStrategyRequest() (request *ModifyOutbandStrategyRequest) {
	request = &ModifyOutbandStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "ModifyOutbandStrategy")
	return
}

func NewModifyOutbandStrategyResponse() (response *ModifyOutbandStrategyResponse) {
	response = &ModifyOutbandStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自定义带外密码策略
func (c *Client) ModifyOutbandStrategy(request *ModifyOutbandStrategyRequest) (response *ModifyOutbandStrategyResponse, err error) {
	if request == nil {
		request = NewModifyOutbandStrategyRequest()
	}
	response = NewModifyOutbandStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryServerStatusRequest() (request *QueryServerStatusRequest) {
	request = &QueryServerStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "QueryServerStatus")
	return
}

func NewQueryServerStatusResponse() (response *QueryServerStatusResponse) {
	response = &QueryServerStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务接口
func (c *Client) QueryServerStatus(request *QueryServerStatusRequest) (response *QueryServerStatusResponse, err error) {
	if request == nil {
		request = NewQueryServerStatusRequest()
	}
	response = NewQueryServerStatusResponse()
	err = c.Send(request, response)
	return
}

func NewAllocateVMVirtualIPRequest() (request *AllocateVMVirtualIPRequest) {
	request = &AllocateVMVirtualIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("dcos", APIVersion, "AllocateVMVirtualIP")
	return
}

func NewAllocateVMVirtualIPResponse() (response *AllocateVMVirtualIPResponse) {
	response = &AllocateVMVirtualIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分配VM虚拟内网IP
func (c *Client) AllocateVMVirtualIP(request *AllocateVMVirtualIPRequest) (response *AllocateVMVirtualIPResponse, err error) {
	if request == nil {
		request = NewAllocateVMVirtualIPRequest()
	}
	response = NewAllocateVMVirtualIPResponse()
	err = c.Send(request, response)
	return
}
