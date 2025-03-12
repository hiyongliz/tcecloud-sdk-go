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

package v20180125

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2018-01-25"

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

func NewWafSetAddCustomRuleRequest() (request *WafSetAddCustomRuleRequest) {
	request = &WafSetAddCustomRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetAddCustomRule")
	return
}

func NewWafSetAddCustomRuleResponse() (response *WafSetAddCustomRuleResponse) {
	response = &WafSetAddCustomRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加自定义规则
func (c *Client) WafSetAddCustomRule(request *WafSetAddCustomRuleRequest) (response *WafSetAddCustomRuleResponse, err error) {
	if request == nil {
		request = NewWafSetAddCustomRuleRequest()
	}
	response = NewWafSetAddCustomRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetBwIPRequest() (request *WafGetBwIPRequest) {
	request = &WafGetBwIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetBwIP")
	return
}

func NewWafGetBwIPResponse() (response *WafGetBwIPResponse) {
	response = &WafGetBwIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取IP黑白名单
func (c *Client) WafGetBwIP(request *WafGetBwIPRequest) (response *WafGetBwIPResponse, err error) {
	if request == nil {
		request = NewWafGetBwIPRequest()
	}
	response = NewWafGetBwIPResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetVipRequest() (request *WafGetVipRequest) {
	request = &WafGetVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetVip")
	return
}

func NewWafGetVipResponse() (response *WafGetVipResponse) {
	response = &WafGetVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取vip列表
func (c *Client) WafGetVip(request *WafGetVipRequest) (response *WafGetVipResponse, err error) {
	if request == nil {
		request = NewWafGetVipRequest()
	}
	response = NewWafGetVipResponse()
	err = c.Send(request, response)
	return
}

func NewWafVipUpdateRsRequest() (request *WafVipUpdateRsRequest) {
	request = &WafVipUpdateRsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafVipUpdateRs")
	return
}

func NewWafVipUpdateRsResponse() (response *WafVipUpdateRsResponse) {
	response = &WafVipUpdateRsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑VIP
func (c *Client) WafVipUpdateRs(request *WafVipUpdateRsRequest) (response *WafVipUpdateRsResponse, err error) {
	if request == nil {
		request = NewWafVipUpdateRsRequest()
	}
	response = NewWafVipUpdateRsResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetUpgredeProgressRequest() (request *WafGetUpgredeProgressRequest) {
	request = &WafGetUpgredeProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetUpgredeProgress")
	return
}

func NewWafGetUpgredeProgressResponse() (response *WafGetUpgredeProgressResponse) {
	response = &WafGetUpgredeProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得手动上传进度
func (c *Client) WafGetUpgredeProgress(request *WafGetUpgredeProgressRequest) (response *WafGetUpgredeProgressResponse, err error) {
	if request == nil {
		request = NewWafGetUpgredeProgressRequest()
	}
	response = NewWafGetUpgredeProgressResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetProductUsagesRequest() (request *WafGetProductUsagesRequest) {
	request = &WafGetProductUsagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetProductUsages")
	return
}

func NewWafGetProductUsagesResponse() (response *WafGetProductUsagesResponse) {
	response = &WafGetProductUsagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// license获取waf产品用量
func (c *Client) WafGetProductUsages(request *WafGetProductUsagesRequest) (response *WafGetProductUsagesResponse, err error) {
	if request == nil {
		request = NewWafGetProductUsagesRequest()
	}
	response = NewWafGetProductUsagesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGlobalWhiteRuleRequest() (request *ModifyGlobalWhiteRuleRequest) {
	request = &ModifyGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyGlobalWhiteRule")
	return
}

func NewModifyGlobalWhiteRuleResponse() (response *ModifyGlobalWhiteRuleResponse) {
	response = &ModifyGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：修改规则引擎全局白名单
func (c *Client) ModifyGlobalWhiteRule(request *ModifyGlobalWhiteRuleRequest) (response *ModifyGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewModifyGlobalWhiteRuleRequest()
	}
	response = NewModifyGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewTCERuleUploadRequest() (request *TCERuleUploadRequest) {
	request = &TCERuleUploadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "TCERuleUpload")
	return
}

func NewTCERuleUploadResponse() (response *TCERuleUploadResponse) {
	response = &TCERuleUploadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传tiga规则库信息到mysql
func (c *Client) TCERuleUpload(request *TCERuleUploadRequest) (response *TCERuleUploadResponse, err error) {
	if request == nil {
		request = NewTCERuleUploadRequest()
	}
	response = NewTCERuleUploadResponse()
	err = c.Send(request, response)
	return
}

func NewEditOpRuleUpdateLogRequest() (request *EditOpRuleUpdateLogRequest) {
	request = &EditOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EditOpRuleUpdateLog")
	return
}

func NewEditOpRuleUpdateLogResponse() (response *EditOpRuleUpdateLogResponse) {
	response = &EditOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：编辑规则更新日志
func (c *Client) EditOpRuleUpdateLog(request *EditOpRuleUpdateLogRequest) (response *EditOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewEditOpRuleUpdateLogRequest()
	}
	response = NewEditOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetRuleUpgradeRecordRequest() (request *WafGetRuleUpgradeRecordRequest) {
	request = &WafGetRuleUpgradeRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetRuleUpgradeRecord")
	return
}

func NewWafGetRuleUpgradeRecordResponse() (response *WafGetRuleUpgradeRecordResponse) {
	response = &WafGetRuleUpgradeRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得版本升级记录
func (c *Client) WafGetRuleUpgradeRecord(request *WafGetRuleUpgradeRecordRequest) (response *WafGetRuleUpgradeRecordResponse, err error) {
	if request == nil {
		request = NewWafGetRuleUpgradeRecordRequest()
	}
	response = NewWafGetRuleUpgradeRecordResponse()
	err = c.Send(request, response)
	return
}

func NewGetTCERuleCosKeyRequest() (request *GetTCERuleCosKeyRequest) {
	request = &GetTCERuleCosKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetTCERuleCosKey")
	return
}

func NewGetTCERuleCosKeyResponse() (response *GetTCERuleCosKeyResponse) {
	response = &GetTCERuleCosKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取cos临时密钥，用于上传tiga规则库
func (c *Client) GetTCERuleCosKey(request *GetTCERuleCosKeyRequest) (response *GetTCERuleCosKeyResponse, err error) {
	if request == nil {
		request = NewGetTCERuleCosKeyRequest()
	}
	response = NewGetTCERuleCosKeyResponse()
	err = c.Send(request, response)
	return
}

func NewTCERuleRollbackRequest() (request *TCERuleRollbackRequest) {
	request = &TCERuleRollbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "TCERuleRollback")
	return
}

func NewTCERuleRollbackResponse() (response *TCERuleRollbackResponse) {
	response = &TCERuleRollbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚tiga规则库
func (c *Client) TCERuleRollback(request *TCERuleRollbackRequest) (response *TCERuleRollbackResponse, err error) {
	if request == nil {
		request = NewTCERuleRollbackRequest()
	}
	response = NewTCERuleRollbackResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetNgWafVipConfRequest() (request *WafSetNgWafVipConfRequest) {
	request = &WafSetNgWafVipConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetNgWafVipConf")
	return
}

func NewWafSetNgWafVipConfResponse() (response *WafSetNgWafVipConfResponse) {
	response = &WafSetNgWafVipConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加CLBWAF集群管理
func (c *Client) WafSetNgWafVipConf(request *WafSetNgWafVipConfRequest) (response *WafSetNgWafVipConfResponse, err error) {
	if request == nil {
		request = NewWafSetNgWafVipConfRequest()
	}
	response = NewWafSetNgWafVipConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomWhiteRuleRequest() (request *DescribeCustomWhiteRuleRequest) {
	request = &DescribeCustomWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeCustomWhiteRule")
	return
}

func NewDescribeCustomWhiteRuleResponse() (response *DescribeCustomWhiteRuleResponse) {
	response = &DescribeCustomWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 精准白名单查询
func (c *Client) DescribeCustomWhiteRule(request *DescribeCustomWhiteRuleRequest) (response *DescribeCustomWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCustomWhiteRuleRequest()
	}
	response = NewDescribeCustomWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOpSignatureRuleRequest() (request *ModifyOpSignatureRuleRequest) {
	request = &ModifyOpSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyOpSignatureRule")
	return
}

func NewModifyOpSignatureRuleResponse() (response *ModifyOpSignatureRuleResponse) {
	response = &ModifyOpSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：修改规则
func (c *Client) ModifyOpSignatureRule(request *ModifyOpSignatureRuleRequest) (response *ModifyOpSignatureRuleResponse, err error) {
	if request == nil {
		request = NewModifyOpSignatureRuleRequest()
	}
	response = NewModifyOpSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafUpsertSysLogRequest() (request *WafUpsertSysLogRequest) {
	request = &WafUpsertSysLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafUpsertSysLog")
	return
}

func NewWafUpsertSysLogResponse() (response *WafUpsertSysLogResponse) {
	response = &WafUpsertSysLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置syslog
func (c *Client) WafUpsertSysLog(request *WafUpsertSysLogRequest) (response *WafUpsertSysLogResponse, err error) {
	if request == nil {
		request = NewWafUpsertSysLogRequest()
	}
	response = NewWafUpsertSysLogResponse()
	err = c.Send(request, response)
	return
}

func NewAddGlobalWhiteRuleRequest() (request *AddGlobalWhiteRuleRequest) {
	request = &AddGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "AddGlobalWhiteRule")
	return
}

func NewAddGlobalWhiteRuleResponse() (response *AddGlobalWhiteRuleResponse) {
	response = &AddGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：增加规则引擎全局白名单
func (c *Client) AddGlobalWhiteRule(request *AddGlobalWhiteRuleRequest) (response *AddGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewAddGlobalWhiteRuleRequest()
	}
	response = NewAddGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafBanIPRequest() (request *WafBanIPRequest) {
	request = &WafBanIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafBanIP")
	return
}

func NewWafBanIPResponse() (response *WafBanIPResponse) {
	response = &WafBanIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 天幕封禁
func (c *Client) WafBanIP(request *WafBanIPRequest) (response *WafBanIPResponse, err error) {
	if request == nil {
		request = NewWafBanIPRequest()
	}
	response = NewWafBanIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpUserWhiteRuleRequest() (request *DescribeOpUserWhiteRuleRequest) {
	request = &DescribeOpUserWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpUserWhiteRule")
	return
}

func NewDescribeOpUserWhiteRuleResponse() (response *DescribeOpUserWhiteRuleResponse) {
	response = &DescribeOpUserWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台接口：获取用户规则白名单列表
func (c *Client) DescribeOpUserWhiteRule(request *DescribeOpUserWhiteRuleRequest) (response *DescribeOpUserWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpUserWhiteRuleRequest()
	}
	response = NewDescribeOpUserWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetHostsRequest() (request *WafGetHostsRequest) {
	request = &WafGetHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetHosts")
	return
}

func NewWafGetHostsResponse() (response *WafGetHostsResponse) {
	response = &WafGetHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名列表
func (c *Client) WafGetHosts(request *WafGetHostsRequest) (response *WafGetHostsResponse, err error) {
	if request == nil {
		request = NewWafGetHostsRequest()
	}
	response = NewWafGetHostsResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetRulesRequest() (request *WafGetRulesRequest) {
	request = &WafGetRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetRules")
	return
}

func NewWafGetRulesResponse() (response *WafGetRulesResponse) {
	response = &WafGetRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 客户自定义规则管理
func (c *Client) WafGetRules(request *WafGetRulesRequest) (response *WafGetRulesResponse, err error) {
	if request == nil {
		request = NewWafGetRulesRequest()
	}
	response = NewWafGetRulesResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetRuleUpgradeConfRequest() (request *WafGetRuleUpgradeConfRequest) {
	request = &WafGetRuleUpgradeConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetRuleUpgradeConf")
	return
}

func NewWafGetRuleUpgradeConfResponse() (response *WafGetRuleUpgradeConfResponse) {
	response = &WafGetRuleUpgradeConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得规则自动升级配置
func (c *Client) WafGetRuleUpgradeConf(request *WafGetRuleUpgradeConfRequest) (response *WafGetRuleUpgradeConfResponse, err error) {
	if request == nil {
		request = NewWafGetRuleUpgradeConfRequest()
	}
	response = NewWafGetRuleUpgradeConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpUserSignaturePolicyRequest() (request *DescribeOpUserSignaturePolicyRequest) {
	request = &DescribeOpUserSignaturePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpUserSignaturePolicy")
	return
}

func NewDescribeOpUserSignaturePolicyResponse() (response *DescribeOpUserSignaturePolicyResponse) {
	response = &DescribeOpUserSignaturePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取永辉信息
func (c *Client) DescribeOpUserSignaturePolicy(request *DescribeOpUserSignaturePolicyRequest) (response *DescribeOpUserSignaturePolicyResponse, err error) {
	if request == nil {
		request = NewDescribeOpUserSignaturePolicyRequest()
	}
	response = NewDescribeOpUserSignaturePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewWafDelNgWafVipConfRequest() (request *WafDelNgWafVipConfRequest) {
	request = &WafDelNgWafVipConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafDelNgWafVipConf")
	return
}

func NewWafDelNgWafVipConfResponse() (response *WafDelNgWafVipConfResponse) {
	response = &WafDelNgWafVipConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除CLBWAF集群管理
func (c *Client) WafDelNgWafVipConf(request *WafDelNgWafVipConfRequest) (response *WafDelNgWafVipConfResponse, err error) {
	if request == nil {
		request = NewWafDelNgWafVipConfRequest()
	}
	response = NewWafDelNgWafVipConfResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetConfigSwitchInfoRequest() (request *WafGetConfigSwitchInfoRequest) {
	request = &WafGetConfigSwitchInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetConfigSwitchInfo")
	return
}

func NewWafGetConfigSwitchInfoResponse() (response *WafGetConfigSwitchInfoResponse) {
	response = &WafGetConfigSwitchInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 防护默认开关设置
func (c *Client) WafGetConfigSwitchInfo(request *WafGetConfigSwitchInfoRequest) (response *WafGetConfigSwitchInfoResponse, err error) {
	if request == nil {
		request = NewWafGetConfigSwitchInfoRequest()
	}
	response = NewWafGetConfigSwitchInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpMainClassRequest() (request *DescribeOpMainClassRequest) {
	request = &DescribeOpMainClassRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpMainClass")
	return
}

func NewDescribeOpMainClassResponse() (response *DescribeOpMainClassResponse) {
	response = &DescribeOpMainClassResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取主类及子类信息
func (c *Client) DescribeOpMainClass(request *DescribeOpMainClassRequest) (response *DescribeOpMainClassResponse, err error) {
	if request == nil {
		request = NewDescribeOpMainClassRequest()
	}
	response = NewDescribeOpMainClassResponse()
	err = c.Send(request, response)
	return
}

func NewWafUpsertNonstandardPortRequest() (request *WafUpsertNonstandardPortRequest) {
	request = &WafUpsertNonstandardPortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafUpsertNonstandardPort")
	return
}

func NewWafUpsertNonstandardPortResponse() (response *WafUpsertNonstandardPortResponse) {
	response = &WafUpsertNonstandardPortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf新增或更新非标端口
func (c *Client) WafUpsertNonstandardPort(request *WafUpsertNonstandardPortRequest) (response *WafUpsertNonstandardPortResponse, err error) {
	if request == nil {
		request = NewWafUpsertNonstandardPortRequest()
	}
	response = NewWafUpsertNonstandardPortResponse()
	err = c.Send(request, response)
	return
}

func NewTCERuleUpgradeRequest() (request *TCERuleUpgradeRequest) {
	request = &TCERuleUpgradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "TCERuleUpgrade")
	return
}

func NewTCERuleUpgradeResponse() (response *TCERuleUpgradeResponse) {
	response = &TCERuleUpgradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动tiga规则库升级任务
func (c *Client) TCERuleUpgrade(request *TCERuleUpgradeRequest) (response *TCERuleUpgradeResponse, err error) {
	if request == nil {
		request = NewTCERuleUpgradeRequest()
	}
	response = NewTCERuleUpgradeResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetLatestRuleVersionRequest() (request *WafGetLatestRuleVersionRequest) {
	request = &WafGetLatestRuleVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetLatestRuleVersion")
	return
}

func NewWafGetLatestRuleVersionResponse() (response *WafGetLatestRuleVersionResponse) {
	response = &WafGetLatestRuleVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得最新规则版本
func (c *Client) WafGetLatestRuleVersion(request *WafGetLatestRuleVersionRequest) (response *WafGetLatestRuleVersionResponse, err error) {
	if request == nil {
		request = NewWafGetLatestRuleVersionRequest()
	}
	response = NewWafGetLatestRuleVersionResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetDomainEngineRequest() (request *WafSetDomainEngineRequest) {
	request = &WafSetDomainEngineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetDomainEngine")
	return
}

func NewWafSetDomainEngineResponse() (response *WafSetDomainEngineResponse) {
	response = &WafSetDomainEngineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按域名设置引擎类型
func (c *Client) WafSetDomainEngine(request *WafSetDomainEngineRequest) (response *WafSetDomainEngineResponse, err error) {
	if request == nil {
		request = NewWafSetDomainEngineRequest()
	}
	response = NewWafSetDomainEngineResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostStatusRequest() (request *ModifyHostStatusRequest) {
	request = &ModifyHostStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostStatus")
	return
}

func NewModifyHostStatusResponse() (response *ModifyHostStatusResponse) {
	response = &ModifyHostStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb防护域名开关切换
func (c *Client) ModifyHostStatus(request *ModifyHostStatusRequest) (response *ModifyHostStatusResponse, err error) {
	if request == nil {
		request = NewModifyHostStatusRequest()
	}
	response = NewModifyHostStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpSignatureRuleRequest() (request *DescribeOpSignatureRuleRequest) {
	request = &DescribeOpSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpSignatureRule")
	return
}

func NewDescribeOpSignatureRuleResponse() (response *DescribeOpSignatureRuleResponse) {
	response = &DescribeOpSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口，获取全量特征规则列表
func (c *Client) DescribeOpSignatureRule(request *DescribeOpSignatureRuleRequest) (response *DescribeOpSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpSignatureRuleRequest()
	}
	response = NewDescribeOpSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafAddVipRequest() (request *WafAddVipRequest) {
	request = &WafAddVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafAddVip")
	return
}

func NewWafAddVipResponse() (response *WafAddVipResponse) {
	response = &WafAddVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加vip
func (c *Client) WafAddVip(request *WafAddVipRequest) (response *WafAddVipResponse, err error) {
	if request == nil {
		request = NewWafAddVipRequest()
	}
	response = NewWafAddVipResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetCheckDomainRequest() (request *WafGetCheckDomainRequest) {
	request = &WafGetCheckDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetCheckDomain")
	return
}

func NewWafGetCheckDomainResponse() (response *WafGetCheckDomainResponse) {
	response = &WafGetCheckDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查域名是否可用
func (c *Client) WafGetCheckDomain(request *WafGetCheckDomainRequest) (response *WafGetCheckDomainResponse, err error) {
	if request == nil {
		request = NewWafGetCheckDomainRequest()
	}
	response = NewWafGetCheckDomainResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetCustomPayloadRequest() (request *WafSetCustomPayloadRequest) {
	request = &WafSetCustomPayloadRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetCustomPayload")
	return
}

func NewWafSetCustomPayloadResponse() (response *WafSetCustomPayloadResponse) {
	response = &WafSetCustomPayloadResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加误报漏报荷载
func (c *Client) WafSetCustomPayload(request *WafSetCustomPayloadRequest) (response *WafSetCustomPayloadResponse, err error) {
	if request == nil {
		request = NewWafSetCustomPayloadRequest()
	}
	response = NewWafSetCustomPayloadResponse()
	err = c.Send(request, response)
	return
}

func NewTCERuleDeleteRequest() (request *TCERuleDeleteRequest) {
	request = &TCERuleDeleteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "TCERuleDelete")
	return
}

func NewTCERuleDeleteResponse() (response *TCERuleDeleteResponse) {
	response = &TCERuleDeleteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除tiga规则库
func (c *Client) TCERuleDelete(request *TCERuleDeleteRequest) (response *TCERuleDeleteResponse, err error) {
	if request == nil {
		request = NewTCERuleDeleteRequest()
	}
	response = NewTCERuleDeleteResponse()
	err = c.Send(request, response)
	return
}

func NewWafAddTestClientRequest() (request *WafAddTestClientRequest) {
	request = &WafAddTestClientRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafAddTestClient")
	return
}

func NewWafAddTestClientResponse() (response *WafAddTestClientResponse) {
	response = &WafAddTestClientResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加测试用户
func (c *Client) WafAddTestClient(request *WafAddTestClientRequest) (response *WafAddTestClientResponse, err error) {
	if request == nil {
		request = NewWafAddTestClientRequest()
	}
	response = NewWafAddTestClientResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetExportAttackLogRequest() (request *WafGetExportAttackLogRequest) {
	request = &WafGetExportAttackLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetExportAttackLog")
	return
}

func NewWafGetExportAttackLogResponse() (response *WafGetExportAttackLogResponse) {
	response = &WafGetExportAttackLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出攻击日志
func (c *Client) WafGetExportAttackLog(request *WafGetExportAttackLogRequest) (response *WafGetExportAttackLogResponse, err error) {
	if request == nil {
		request = NewWafGetExportAttackLogRequest()
	}
	response = NewWafGetExportAttackLogResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetInstancesRequest() (request *WafGetInstancesRequest) {
	request = &WafGetInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetInstances")
	return
}

func NewWafGetInstancesResponse() (response *WafGetInstancesResponse) {
	response = &WafGetInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例列表
func (c *Client) WafGetInstances(request *WafGetInstancesRequest) (response *WafGetInstancesResponse, err error) {
	if request == nil {
		request = NewWafGetInstancesRequest()
	}
	response = NewWafGetInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewEditOpRuleUpdateLogStatusRequest() (request *EditOpRuleUpdateLogStatusRequest) {
	request = &EditOpRuleUpdateLogStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "EditOpRuleUpdateLogStatus")
	return
}

func NewEditOpRuleUpdateLogStatusResponse() (response *EditOpRuleUpdateLogStatusResponse) {
	response = &EditOpRuleUpdateLogStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：编辑规则更新日志的发布状态
func (c *Client) EditOpRuleUpdateLogStatus(request *EditOpRuleUpdateLogStatusRequest) (response *EditOpRuleUpdateLogStatusResponse, err error) {
	if request == nil {
		request = NewEditOpRuleUpdateLogStatusRequest()
	}
	response = NewEditOpRuleUpdateLogStatusResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetEnableSysLogRequest() (request *WafSetEnableSysLogRequest) {
	request = &WafSetEnableSysLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetEnableSysLog")
	return
}

func NewWafSetEnableSysLogResponse() (response *WafSetEnableSysLogResponse) {
	response = &WafSetEnableSysLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开关syslog
func (c *Client) WafSetEnableSysLog(request *WafSetEnableSysLogRequest) (response *WafSetEnableSysLogResponse, err error) {
	if request == nil {
		request = NewWafSetEnableSysLogRequest()
	}
	response = NewWafSetEnableSysLogResponse()
	err = c.Send(request, response)
	return
}

func NewGetNgWafHostsRequest() (request *GetNgWafHostsRequest) {
	request = &GetNgWafHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "GetNgWafHosts")
	return
}

func NewGetNgWafHostsResponse() (response *GetNgWafHostsResponse) {
	response = &GetNgWafHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CLBWAF域名管理
func (c *Client) GetNgWafHosts(request *GetNgWafHostsRequest) (response *GetNgWafHostsResponse, err error) {
	if request == nil {
		request = NewGetNgWafHostsRequest()
	}
	response = NewGetNgWafHostsResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetCustomPayloadsRequest() (request *WafGetCustomPayloadsRequest) {
	request = &WafGetCustomPayloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetCustomPayloads")
	return
}

func NewWafGetCustomPayloadsResponse() (response *WafGetCustomPayloadsResponse) {
	response = &WafGetCustomPayloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取误报漏报荷载列表
func (c *Client) WafGetCustomPayloads(request *WafGetCustomPayloadsRequest) (response *WafGetCustomPayloadsResponse, err error) {
	if request == nil {
		request = NewWafGetCustomPayloadsRequest()
	}
	response = NewWafGetCustomPayloadsResponse()
	err = c.Send(request, response)
	return
}

func NewWafDeleteSpartaProtectionRequest() (request *WafDeleteSpartaProtectionRequest) {
	request = &WafDeleteSpartaProtectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafDeleteSpartaProtection")
	return
}

func NewWafDeleteSpartaProtectionResponse() (response *WafDeleteSpartaProtectionResponse) {
	response = &WafDeleteSpartaProtectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除域名host_host
func (c *Client) WafDeleteSpartaProtection(request *WafDeleteSpartaProtectionRequest) (response *WafDeleteSpartaProtectionResponse, err error) {
	if request == nil {
		request = NewWafDeleteSpartaProtectionRequest()
	}
	response = NewWafDeleteSpartaProtectionResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetCustomConfigRequest() (request *WafSetCustomConfigRequest) {
	request = &WafSetCustomConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetCustomConfig")
	return
}

func NewWafSetCustomConfigResponse() (response *WafSetCustomConfigResponse) {
	response = &WafSetCustomConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// xff_status、ulimit_args配置设置
func (c *Client) WafSetCustomConfig(request *WafSetCustomConfigRequest) (response *WafSetCustomConfigResponse, err error) {
	if request == nil {
		request = NewWafSetCustomConfigRequest()
	}
	response = NewWafSetCustomConfigResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetSwitchCustomRuleStatusRequest() (request *WafSetSwitchCustomRuleStatusRequest) {
	request = &WafSetSwitchCustomRuleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetSwitchCustomRuleStatus")
	return
}

func NewWafSetSwitchCustomRuleStatusResponse() (response *WafSetSwitchCustomRuleStatusResponse) {
	response = &WafSetSwitchCustomRuleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 切换自定义规则状态
func (c *Client) WafSetSwitchCustomRuleStatus(request *WafSetSwitchCustomRuleStatusRequest) (response *WafSetSwitchCustomRuleStatusResponse, err error) {
	if request == nil {
		request = NewWafSetSwitchCustomRuleStatusRequest()
	}
	response = NewWafSetSwitchCustomRuleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetDomainEngineTypeRequest() (request *WafGetDomainEngineTypeRequest) {
	request = &WafGetDomainEngineTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetDomainEngineType")
	return
}

func NewWafGetDomainEngineTypeResponse() (response *WafGetDomainEngineTypeResponse) {
	response = &WafGetDomainEngineTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名使用的规则引擎类型
func (c *Client) WafGetDomainEngineType(request *WafGetDomainEngineTypeRequest) (response *WafGetDomainEngineTypeResponse, err error) {
	if request == nil {
		request = NewWafGetDomainEngineTypeRequest()
	}
	response = NewWafGetDomainEngineTypeResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetFreqRulesRequest() (request *WafGetFreqRulesRequest) {
	request = &WafGetFreqRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetFreqRules")
	return
}

func NewWafGetFreqRulesResponse() (response *WafGetFreqRulesResponse) {
	response = &WafGetFreqRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CC规则
func (c *Client) WafGetFreqRules(request *WafGetFreqRulesRequest) (response *WafGetFreqRulesResponse, err error) {
	if request == nil {
		request = NewWafGetFreqRulesRequest()
	}
	response = NewWafGetFreqRulesResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetLicenseListRequest() (request *WafGetLicenseListRequest) {
	request = &WafGetLicenseListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetLicenseList")
	return
}

func NewWafGetLicenseListResponse() (response *WafGetLicenseListResponse) {
	response = &WafGetLicenseListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取license列表
func (c *Client) WafGetLicenseList(request *WafGetLicenseListRequest) (response *WafGetLicenseListResponse, err error) {
	if request == nil {
		request = NewWafGetLicenseListRequest()
	}
	response = NewWafGetLicenseListResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetSysLogRequest() (request *WafGetSysLogRequest) {
	request = &WafGetSysLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetSysLog")
	return
}

func NewWafGetSysLogResponse() (response *WafGetSysLogResponse) {
	response = &WafGetSysLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取syslog配置
func (c *Client) WafGetSysLog(request *WafGetSysLogRequest) (response *WafGetSysLogResponse, err error) {
	if request == nil {
		request = NewWafGetSysLogRequest()
	}
	response = NewWafGetSysLogResponse()
	err = c.Send(request, response)
	return
}

func NewWafDeleteInstancesRequest() (request *WafDeleteInstancesRequest) {
	request = &WafDeleteInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafDeleteInstances")
	return
}

func NewWafDeleteInstancesResponse() (response *WafDeleteInstancesResponse) {
	response = &WafDeleteInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除实例
func (c *Client) WafDeleteInstances(request *WafDeleteInstancesRequest) (response *WafDeleteInstancesResponse, err error) {
	if request == nil {
		request = NewWafDeleteInstancesRequest()
	}
	response = NewWafDeleteInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetExportAttackLogJobsRequest() (request *WafGetExportAttackLogJobsRequest) {
	request = &WafGetExportAttackLogJobsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetExportAttackLogJobs")
	return
}

func NewWafGetExportAttackLogJobsResponse() (response *WafGetExportAttackLogJobsResponse) {
	response = &WafGetExportAttackLogJobsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得日志导出任务列表
func (c *Client) WafGetExportAttackLogJobs(request *WafGetExportAttackLogJobsRequest) (response *WafGetExportAttackLogJobsResponse, err error) {
	if request == nil {
		request = NewWafGetExportAttackLogJobsRequest()
	}
	response = NewWafGetExportAttackLogJobsResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetCacheRulesRequest() (request *WafGetCacheRulesRequest) {
	request = &WafGetCacheRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetCacheRules")
	return
}

func NewWafGetCacheRulesResponse() (response *WafGetCacheRulesResponse) {
	response = &WafGetCacheRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取防篡改url
func (c *Client) WafGetCacheRules(request *WafGetCacheRulesRequest) (response *WafGetCacheRulesResponse, err error) {
	if request == nil {
		request = NewWafGetCacheRulesRequest()
	}
	response = NewWafGetCacheRulesResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetWafDomainLoadbalancerRequest() (request *WafGetWafDomainLoadbalancerRequest) {
	request = &WafGetWafDomainLoadbalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetWafDomainLoadbalancer")
	return
}

func NewWafGetWafDomainLoadbalancerResponse() (response *WafGetWafDomainLoadbalancerResponse) {
	response = &WafGetWafDomainLoadbalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CLBWAF监听器管理
func (c *Client) WafGetWafDomainLoadbalancer(request *WafGetWafDomainLoadbalancerRequest) (response *WafGetWafDomainLoadbalancerResponse, err error) {
	if request == nil {
		request = NewWafGetWafDomainLoadbalancerRequest()
	}
	response = NewWafGetWafDomainLoadbalancerResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetEditHostRequest() (request *WafSetEditHostRequest) {
	request = &WafSetEditHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetEditHost")
	return
}

func NewWafSetEditHostResponse() (response *WafSetEditHostResponse) {
	response = &WafSetEditHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ipv6添加域名
func (c *Client) WafSetEditHost(request *WafSetEditHostRequest) (response *WafSetEditHostResponse, err error) {
	if request == nil {
		request = NewWafSetEditHostRequest()
	}
	response = NewWafSetEditHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpUserSignatureRuleRequest() (request *DescribeOpUserSignatureRuleRequest) {
	request = &DescribeOpUserSignatureRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpUserSignatureRule")
	return
}

func NewDescribeOpUserSignatureRuleResponse() (response *DescribeOpUserSignatureRuleResponse) {
	response = &DescribeOpUserSignatureRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为运营后台获取用户规则
func (c *Client) DescribeOpUserSignatureRule(request *DescribeOpUserSignatureRuleRequest) (response *DescribeOpUserSignatureRuleResponse, err error) {
	if request == nil {
		request = NewDescribeOpUserSignatureRuleRequest()
	}
	response = NewDescribeOpUserSignatureRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetAttackDetailRequest() (request *WafGetAttackDetailRequest) {
	request = &WafGetAttackDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetAttackDetail")
	return
}

func NewWafGetAttackDetailResponse() (response *WafGetAttackDetailResponse) {
	response = &WafGetAttackDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取攻击数据
func (c *Client) WafGetAttackDetail(request *WafGetAttackDetailRequest) (response *WafGetAttackDetailResponse, err error) {
	if request == nil {
		request = NewWafGetAttackDetailRequest()
	}
	response = NewWafGetAttackDetailResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetQueryHostRequest() (request *WafGetQueryHostRequest) {
	request = &WafGetQueryHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetQueryHost")
	return
}

func NewWafGetQueryHostResponse() (response *WafGetQueryHostResponse) {
	response = &WafGetQueryHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得clb域名列表
func (c *Client) WafGetQueryHost(request *WafGetQueryHostRequest) (response *WafGetQueryHostResponse, err error) {
	if request == nil {
		request = NewWafGetQueryHostRequest()
	}
	response = NewWafGetQueryHostResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetSystemVersionRequest() (request *WafGetSystemVersionRequest) {
	request = &WafGetSystemVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetSystemVersion")
	return
}

func NewWafGetSystemVersionResponse() (response *WafGetSystemVersionResponse) {
	response = &WafGetSystemVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得系统版本
func (c *Client) WafGetSystemVersion(request *WafGetSystemVersionRequest) (response *WafGetSystemVersionResponse, err error) {
	if request == nil {
		request = NewWafGetSystemVersionRequest()
	}
	response = NewWafGetSystemVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTCERuleRequest() (request *DescribeTCERuleRequest) {
	request = &DescribeTCERuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeTCERule")
	return
}

func NewDescribeTCERuleResponse() (response *DescribeTCERuleResponse) {
	response = &DescribeTCERuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示tiga规则库详细列表
func (c *Client) DescribeTCERule(request *DescribeTCERuleRequest) (response *DescribeTCERuleResponse, err error) {
	if request == nil {
		request = NewDescribeTCERuleRequest()
	}
	response = NewDescribeTCERuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetWafHostsRequest() (request *WafGetWafHostsRequest) {
	request = &WafGetWafHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetWafHosts")
	return
}

func NewWafGetWafHostsResponse() (response *WafGetWafHostsResponse) {
	response = &WafGetWafHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// waf地域列表
func (c *Client) WafGetWafHosts(request *WafGetWafHostsRequest) (response *WafGetWafHostsResponse, err error) {
	if request == nil {
		request = NewWafGetWafHostsRequest()
	}
	response = NewWafGetWafHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGlobalWhiteRuleRequest() (request *DescribeGlobalWhiteRuleRequest) {
	request = &DescribeGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeGlobalWhiteRule")
	return
}

func NewDescribeGlobalWhiteRuleResponse() (response *DescribeGlobalWhiteRuleResponse) {
	response = &DescribeGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：获取全局白名单
func (c *Client) DescribeGlobalWhiteRule(request *DescribeGlobalWhiteRuleRequest) (response *DescribeGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeGlobalWhiteRuleRequest()
	}
	response = NewDescribeGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyHostModeRequest() (request *ModifyHostModeRequest) {
	request = &ModifyHostModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifyHostMode")
	return
}

func NewModifyHostModeResponse() (response *ModifyHostModeResponse) {
	response = &ModifyHostModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// clb-waf设置防护域名防护状态
func (c *Client) ModifyHostMode(request *ModifyHostModeRequest) (response *ModifyHostModeResponse, err error) {
	if request == nil {
		request = NewModifyHostModeRequest()
	}
	response = NewModifyHostModeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOpRuleUpdateLogRequest() (request *DeleteOpRuleUpdateLogRequest) {
	request = &DeleteOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteOpRuleUpdateLog")
	return
}

func NewDeleteOpRuleUpdateLogResponse() (response *DeleteOpRuleUpdateLogResponse) {
	response = &DeleteOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台接口：删除规则更新日志
func (c *Client) DeleteOpRuleUpdateLog(request *DeleteOpRuleUpdateLogRequest) (response *DeleteOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDeleteOpRuleUpdateLogRequest()
	}
	response = NewDeleteOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetBanIPRequest() (request *WafGetBanIPRequest) {
	request = &WafGetBanIPRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetBanIP")
	return
}

func NewWafGetBanIPResponse() (response *WafGetBanIPResponse) {
	response = &WafGetBanIPResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取天幕封禁参数
func (c *Client) WafGetBanIP(request *WafGetBanIPRequest) (response *WafGetBanIPResponse, err error) {
	if request == nil {
		request = NewWafGetBanIPRequest()
	}
	response = NewWafGetBanIPResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpRuleUpdateLogRequest() (request *DescribeOpRuleUpdateLogRequest) {
	request = &DescribeOpRuleUpdateLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpRuleUpdateLog")
	return
}

func NewDescribeOpRuleUpdateLogResponse() (response *DescribeOpRuleUpdateLogResponse) {
	response = &DescribeOpRuleUpdateLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台：获取特征规则更新动态
func (c *Client) DescribeOpRuleUpdateLog(request *DescribeOpRuleUpdateLogRequest) (response *DescribeOpRuleUpdateLogResponse, err error) {
	if request == nil {
		request = NewDescribeOpRuleUpdateLogRequest()
	}
	response = NewDescribeOpRuleUpdateLogResponse()
	err = c.Send(request, response)
	return
}

func NewWafDelCustomPayloadsRequest() (request *WafDelCustomPayloadsRequest) {
	request = &WafDelCustomPayloadsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafDelCustomPayloads")
	return
}

func NewWafDelCustomPayloadsResponse() (response *WafDelCustomPayloadsResponse) {
	response = &WafDelCustomPayloadsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除误报漏报荷载列表
func (c *Client) WafDelCustomPayloads(request *WafDelCustomPayloadsRequest) (response *WafDelCustomPayloadsResponse, err error) {
	if request == nil {
		request = NewWafDelCustomPayloadsRequest()
	}
	response = NewWafDelCustomPayloadsResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetAppToSuperRequest() (request *WafSetAppToSuperRequest) {
	request = &WafSetAppToSuperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetAppToSuper")
	return
}

func NewWafSetAppToSuperResponse() (response *WafSetAppToSuperResponse) {
	response = &WafSetAppToSuperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置超级管理员
func (c *Client) WafSetAppToSuper(request *WafSetAppToSuperRequest) (response *WafSetAppToSuperResponse, err error) {
	if request == nil {
		request = NewWafSetAppToSuperRequest()
	}
	response = NewWafSetAppToSuperResponse()
	err = c.Send(request, response)
	return
}

func NewCheckGlobalWhiteRuleNameRequest() (request *CheckGlobalWhiteRuleNameRequest) {
	request = &CheckGlobalWhiteRuleNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "CheckGlobalWhiteRuleName")
	return
}

func NewCheckGlobalWhiteRuleNameResponse() (response *CheckGlobalWhiteRuleNameResponse) {
	response = &CheckGlobalWhiteRuleNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台：检测全局白名单名字是否重复
func (c *Client) CheckGlobalWhiteRuleName(request *CheckGlobalWhiteRuleNameRequest) (response *CheckGlobalWhiteRuleNameResponse, err error) {
	if request == nil {
		request = NewCheckGlobalWhiteRuleNameRequest()
	}
	response = NewCheckGlobalWhiteRuleNameResponse()
	err = c.Send(request, response)
	return
}

func NewWafUpdateConfigSwitchRequest() (request *WafUpdateConfigSwitchRequest) {
	request = &WafUpdateConfigSwitchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafUpdateConfigSwitch")
	return
}

func NewWafUpdateConfigSwitchResponse() (response *WafUpdateConfigSwitchResponse) {
	response = &WafUpdateConfigSwitchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改防护默认开关设置
func (c *Client) WafUpdateConfigSwitch(request *WafUpdateConfigSwitchRequest) (response *WafUpdateConfigSwitchResponse, err error) {
	if request == nil {
		request = NewWafUpdateConfigSwitchRequest()
	}
	response = NewWafUpdateConfigSwitchResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGlobalWhiteRuleRequest() (request *DeleteGlobalWhiteRuleRequest) {
	request = &DeleteGlobalWhiteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DeleteGlobalWhiteRule")
	return
}

func NewDeleteGlobalWhiteRuleResponse() (response *DeleteGlobalWhiteRuleResponse) {
	response = &DeleteGlobalWhiteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营后台：删除规则引擎全局白名单
func (c *Client) DeleteGlobalWhiteRule(request *DeleteGlobalWhiteRuleRequest) (response *DeleteGlobalWhiteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteGlobalWhiteRuleRequest()
	}
	response = NewDeleteGlobalWhiteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetClientQpsRequest() (request *WafGetClientQpsRequest) {
	request = &WafGetClientQpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetClientQps")
	return
}

func NewWafGetClientQpsResponse() (response *WafGetClientQpsResponse) {
	response = &WafGetClientQpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客户qps
func (c *Client) WafGetClientQps(request *WafGetClientQpsRequest) (response *WafGetClientQpsResponse, err error) {
	if request == nil {
		request = NewWafGetClientQpsRequest()
	}
	response = NewWafGetClientQpsResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetNgWafHostsRequest() (request *WafGetNgWafHostsRequest) {
	request = &WafGetNgWafHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetNgWafHosts")
	return
}

func NewWafGetNgWafHostsResponse() (response *WafGetNgWafHostsResponse) {
	response = &WafGetNgWafHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CLBWAF域名管理
func (c *Client) WafGetNgWafHosts(request *WafGetNgWafHostsRequest) (response *WafGetNgWafHostsResponse, err error) {
	if request == nil {
		request = NewWafGetNgWafHostsRequest()
	}
	response = NewWafGetNgWafHostsResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetUpgredeRuleByManuallyRequest() (request *WafSetUpgredeRuleByManuallyRequest) {
	request = &WafSetUpgredeRuleByManuallyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetUpgredeRuleByManually")
	return
}

func NewWafSetUpgredeRuleByManuallyResponse() (response *WafSetUpgredeRuleByManuallyResponse) {
	response = &WafSetUpgredeRuleByManuallyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动上传规则版本
func (c *Client) WafSetUpgredeRuleByManually(request *WafSetUpgredeRuleByManuallyRequest) (response *WafSetUpgredeRuleByManuallyResponse, err error) {
	if request == nil {
		request = NewWafSetUpgredeRuleByManuallyRequest()
	}
	response = NewWafSetUpgredeRuleByManuallyResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetCosCredentialsRequest() (request *WafGetCosCredentialsRequest) {
	request = &WafGetCosCredentialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetCosCredentials")
	return
}

func NewWafGetCosCredentialsResponse() (response *WafGetCosCredentialsResponse) {
	response = &WafGetCosCredentialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获得临时授权
func (c *Client) WafGetCosCredentials(request *WafGetCosCredentialsRequest) (response *WafGetCosCredentialsResponse, err error) {
	if request == nil {
		request = NewWafGetCosCredentialsRequest()
	}
	response = NewWafGetCosCredentialsResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetRSRequest() (request *WafGetRSRequest) {
	request = &WafGetRSRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetRS")
	return
}

func NewWafGetRSResponse() (response *WafGetRSResponse) {
	response = &WafGetRSResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取rs列表
func (c *Client) WafGetRS(request *WafGetRSRequest) (response *WafGetRSResponse, err error) {
	if request == nil {
		request = NewWafGetRSRequest()
	}
	response = NewWafGetRSResponse()
	err = c.Send(request, response)
	return
}

func NewWafAddRsRequest() (request *WafAddRsRequest) {
	request = &WafAddRsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafAddRs")
	return
}

func NewWafAddRsResponse() (response *WafAddRsResponse) {
	response = &WafAddRsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加RS
func (c *Client) WafAddRs(request *WafAddRsRequest) (response *WafAddRsResponse, err error) {
	if request == nil {
		request = NewWafAddRsRequest()
	}
	response = NewWafAddRsResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetNonstandardPortRequest() (request *WafGetNonstandardPortRequest) {
	request = &WafGetNonstandardPortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetNonstandardPort")
	return
}

func NewWafGetNonstandardPortResponse() (response *WafGetNonstandardPortResponse) {
	response = &WafGetNonstandardPortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取saaswaf的非标端口
func (c *Client) WafGetNonstandardPort(request *WafGetNonstandardPortRequest) (response *WafGetNonstandardPortResponse, err error) {
	if request == nil {
		request = NewWafGetNonstandardPortRequest()
	}
	response = NewWafGetNonstandardPortResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetRuleUpgradeConfRequest() (request *WafSetRuleUpgradeConfRequest) {
	request = &WafSetRuleUpgradeConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetRuleUpgradeConf")
	return
}

func NewWafSetRuleUpgradeConfResponse() (response *WafSetRuleUpgradeConfResponse) {
	response = &WafSetRuleUpgradeConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存自动升级配置
func (c *Client) WafSetRuleUpgradeConf(request *WafSetRuleUpgradeConfRequest) (response *WafSetRuleUpgradeConfResponse, err error) {
	if request == nil {
		request = NewWafSetRuleUpgradeConfRequest()
	}
	response = NewWafSetRuleUpgradeConfResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOpAttackWhiteRuleNewRequest() (request *DescribeOpAttackWhiteRuleNewRequest) {
	request = &DescribeOpAttackWhiteRuleNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "DescribeOpAttackWhiteRuleNew")
	return
}

func NewDescribeOpAttackWhiteRuleNewResponse() (response *DescribeOpAttackWhiteRuleNewResponse) {
	response = &DescribeOpAttackWhiteRuleNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台接口：获取用户规则白名单列表
func (c *Client) DescribeOpAttackWhiteRuleNew(request *DescribeOpAttackWhiteRuleNewRequest) (response *DescribeOpAttackWhiteRuleNewResponse, err error) {
	if request == nil {
		request = NewDescribeOpAttackWhiteRuleNewRequest()
	}
	response = NewDescribeOpAttackWhiteRuleNewResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetIspsRequest() (request *WafGetIspsRequest) {
	request = &WafGetIspsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetIsps")
	return
}

func NewWafGetIspsResponse() (response *WafGetIspsResponse) {
	response = &WafGetIspsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取运营商列表
func (c *Client) WafGetIsps(request *WafGetIspsRequest) (response *WafGetIspsResponse, err error) {
	if request == nil {
		request = NewWafGetIspsRequest()
	}
	response = NewWafGetIspsResponse()
	err = c.Send(request, response)
	return
}

func NewWafDeleteNonstandardPortRequest() (request *WafDeleteNonstandardPortRequest) {
	request = &WafDeleteNonstandardPortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafDeleteNonstandardPort")
	return
}

func NewWafDeleteNonstandardPortResponse() (response *WafDeleteNonstandardPortResponse) {
	response = &WafDeleteNonstandardPortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除waf非标端口
func (c *Client) WafDeleteNonstandardPort(request *WafDeleteNonstandardPortRequest) (response *WafDeleteNonstandardPortResponse, err error) {
	if request == nil {
		request = NewWafDeleteNonstandardPortRequest()
	}
	response = NewWafDeleteNonstandardPortResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetNgWafVipConfRequest() (request *WafGetNgWafVipConfRequest) {
	request = &WafGetNgWafVipConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetNgWafVipConf")
	return
}

func NewWafGetNgWafVipConfResponse() (response *WafGetNgWafVipConfResponse) {
	response = &WafGetNgWafVipConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取CLBWAF集群管理列表
func (c *Client) WafGetNgWafVipConf(request *WafGetNgWafVipConfRequest) (response *WafGetNgWafVipConfResponse, err error) {
	if request == nil {
		request = NewWafGetNgWafVipConfRequest()
	}
	response = NewWafGetNgWafVipConfResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetQpsRequest() (request *WafGetQpsRequest) {
	request = &WafGetQpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetQps")
	return
}

func NewWafGetQpsResponse() (response *WafGetQpsResponse) {
	response = &WafGetQpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取域名qps
func (c *Client) WafGetQps(request *WafGetQpsRequest) (response *WafGetQpsResponse, err error) {
	if request == nil {
		request = NewWafGetQpsRequest()
	}
	response = NewWafGetQpsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySpartaProtectionModeRequest() (request *ModifySpartaProtectionModeRequest) {
	request = &ModifySpartaProtectionModeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "ModifySpartaProtectionMode")
	return
}

func NewModifySpartaProtectionModeResponse() (response *ModifySpartaProtectionModeResponse) {
	response = &ModifySpartaProtectionModeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置waf防护状态
func (c *Client) ModifySpartaProtectionMode(request *ModifySpartaProtectionModeRequest) (response *ModifySpartaProtectionModeResponse, err error) {
	if request == nil {
		request = NewModifySpartaProtectionModeRequest()
	}
	response = NewModifySpartaProtectionModeResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetNgWafVipRemarkRequest() (request *WafSetNgWafVipRemarkRequest) {
	request = &WafSetNgWafVipRemarkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetNgWafVipRemark")
	return
}

func NewWafSetNgWafVipRemarkResponse() (response *WafSetNgWafVipRemarkResponse) {
	response = &WafSetNgWafVipRemarkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新集群管理备注
func (c *Client) WafSetNgWafVipRemark(request *WafSetNgWafVipRemarkRequest) (response *WafSetNgWafVipRemarkResponse, err error) {
	if request == nil {
		request = NewWafSetNgWafVipRemarkRequest()
	}
	response = NewWafSetNgWafVipRemarkResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetDeleteHostRequest() (request *WafSetDeleteHostRequest) {
	request = &WafSetDeleteHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetDeleteHost")
	return
}

func NewWafSetDeleteHostResponse() (response *WafSetDeleteHostResponse) {
	response = &WafSetDeleteHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除域名host_host
func (c *Client) WafSetDeleteHost(request *WafSetDeleteHostRequest) (response *WafSetDeleteHostResponse, err error) {
	if request == nil {
		request = NewWafSetDeleteHostRequest()
	}
	response = NewWafSetDeleteHostResponse()
	err = c.Send(request, response)
	return
}

func NewWafGetClientsRequest() (request *WafGetClientsRequest) {
	request = &WafGetClientsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafGetClients")
	return
}

func NewWafGetClientsResponse() (response *WafGetClientsResponse) {
	response = &WafGetClientsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取客户列表
func (c *Client) WafGetClients(request *WafGetClientsRequest) (response *WafGetClientsResponse, err error) {
	if request == nil {
		request = NewWafGetClientsRequest()
	}
	response = NewWafGetClientsResponse()
	err = c.Send(request, response)
	return
}

func NewWafSetAddIpv6HostRequest() (request *WafSetAddIpv6HostRequest) {
	request = &WafSetAddIpv6HostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ngwaf", APIVersion, "WafSetAddIpv6Host")
	return
}

func NewWafSetAddIpv6HostResponse() (response *WafSetAddIpv6HostResponse) {
	response = &WafSetAddIpv6HostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ipv6添加域名
func (c *Client) WafSetAddIpv6Host(request *WafSetAddIpv6HostRequest) (response *WafSetAddIpv6HostResponse, err error) {
	if request == nil {
		request = NewWafSetAddIpv6HostRequest()
	}
	response = NewWafSetAddIpv6HostResponse()
	err = c.Send(request, response)
	return
}
