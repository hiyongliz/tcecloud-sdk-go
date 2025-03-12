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

package v20200103

import (
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common"
	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
	"github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/profile"
)

const APIVersion = "2020-01-03"

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

func NewDescribeIspRequest() (request *DescribeIspRequest) {
	request = &DescribeIspRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeIsp")
	return
}

func NewDescribeIspResponse() (response *DescribeIspResponse) {
	response = &DescribeIspResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络运营商
func (c *Client) DescribeIsp(request *DescribeIspRequest) (response *DescribeIspResponse, err error) {
	if request == nil {
		request = NewDescribeIspRequest()
	}
	response = NewDescribeIspResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductHealthStateRequest() (request *QueryProductHealthStateRequest) {
	request = &QueryProductHealthStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "QueryProductHealthState")
	return
}

func NewQueryProductHealthStateResponse() (response *QueryProductHealthStateResponse) {
	response = &QueryProductHealthStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询CLB产品健康状态，结果中应该包盖该产品一个或者多个维度的健康状态和总体健康状态。
func (c *Client) QueryProductHealthState(request *QueryProductHealthStateRequest) (response *QueryProductHealthStateResponse, err error) {
	if request == nil {
		request = NewQueryProductHealthStateRequest()
	}
	response = NewQueryProductHealthStateResponse()
	err = c.Send(request, response)
	return
}

func NewAssignApdLdToSetRequest() (request *AssignApdLdToSetRequest) {
	request = &AssignApdLdToSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "AssignApdLdToSet")
	return
}

func NewAssignApdLdToSetResponse() (response *AssignApdLdToSetResponse) {
	response = &AssignApdLdToSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定LD到Set集群
func (c *Client) AssignApdLdToSet(request *AssignApdLdToSetRequest) (response *AssignApdLdToSetResponse, err error) {
	if request == nil {
		request = NewAssignApdLdToSetRequest()
	}
	response = NewAssignApdLdToSetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApdVipGroupRequest() (request *DeleteApdVipGroupRequest) {
	request = &DeleteApdVipGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteApdVipGroup")
	return
}

func NewDeleteApdVipGroupResponse() (response *DeleteApdVipGroupResponse) {
	response = &DeleteApdVipGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除VIP组
func (c *Client) DeleteApdVipGroup(request *DeleteApdVipGroupRequest) (response *DeleteApdVipGroupResponse, err error) {
	if request == nil {
		request = NewDeleteApdVipGroupRequest()
	}
	response = NewDeleteApdVipGroupResponse()
	err = c.Send(request, response)
	return
}

func NewOpApdStgwSetRequest() (request *OpApdStgwSetRequest) {
	request = &OpApdStgwSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "OpApdStgwSet")
	return
}

func NewOpApdStgwSetResponse() (response *OpApdStgwSetResponse) {
	response = &OpApdStgwSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或删除7层set
func (c *Client) OpApdStgwSet(request *OpApdStgwSetRequest) (response *OpApdStgwSetResponse, err error) {
	if request == nil {
		request = NewOpApdStgwSetRequest()
	}
	response = NewOpApdStgwSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdVipGroupCategoryRequest() (request *DescribeApdVipGroupCategoryRequest) {
	request = &DescribeApdVipGroupCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdVipGroupCategory")
	return
}

func NewDescribeApdVipGroupCategoryResponse() (response *DescribeApdVipGroupCategoryResponse) {
	response = &DescribeApdVipGroupCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VIP组类别详细
func (c *Client) DescribeApdVipGroupCategory(request *DescribeApdVipGroupCategoryRequest) (response *DescribeApdVipGroupCategoryResponse, err error) {
	if request == nil {
		request = NewDescribeApdVipGroupCategoryRequest()
	}
	response = NewDescribeApdVipGroupCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewBindSetLabelRequest() (request *BindSetLabelRequest) {
	request = &BindSetLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "BindSetLabel")
	return
}

func NewBindSetLabelResponse() (response *BindSetLabelResponse) {
	response = &BindSetLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定标签名和tgw集群
func (c *Client) BindSetLabel(request *BindSetLabelRequest) (response *BindSetLabelResponse, err error) {
	if request == nil {
		request = NewBindSetLabelRequest()
	}
	response = NewBindSetLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdLdL7Request() (request *DescribeApdLdL7Request) {
	request = &DescribeApdLdL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdLdL7")
	return
}

func NewDescribeApdLdL7Response() (response *DescribeApdLdL7Response) {
	response = &DescribeApdLdL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询7层LD
func (c *Client) DescribeApdLdL7(request *DescribeApdLdL7Request) (response *DescribeApdLdL7Response, err error) {
	if request == nil {
		request = NewDescribeApdLdL7Request()
	}
	response = NewDescribeApdLdL7Response()
	err = c.Send(request, response)
	return
}

func NewDescribeApdVipGroupRequest() (request *DescribeApdVipGroupRequest) {
	request = &DescribeApdVipGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdVipGroup")
	return
}

func NewDescribeApdVipGroupResponse() (response *DescribeApdVipGroupResponse) {
	response = &DescribeApdVipGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VIP组列表
func (c *Client) DescribeApdVipGroup(request *DescribeApdVipGroupRequest) (response *DescribeApdVipGroupResponse, err error) {
	if request == nil {
		request = NewDescribeApdVipGroupRequest()
	}
	response = NewDescribeApdVipGroupResponse()
	err = c.Send(request, response)
	return
}

func NewProcessApdApplyRequest() (request *ProcessApdApplyRequest) {
	request = &ProcessApdApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ProcessApdApply")
	return
}

func NewProcessApdApplyResponse() (response *ProcessApdApplyResponse) {
	response = &ProcessApdApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实施四层或七层规则申请
func (c *Client) ProcessApdApply(request *ProcessApdApplyRequest) (response *ProcessApdApplyResponse, err error) {
	if request == nil {
		request = NewProcessApdApplyRequest()
	}
	response = NewProcessApdApplyResponse()
	err = c.Send(request, response)
	return
}

func NewSetApdIdcStgwMasterRequest() (request *SetApdIdcStgwMasterRequest) {
	request = &SetApdIdcStgwMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetApdIdcStgwMaster")
	return
}

func NewSetApdIdcStgwMasterResponse() (response *SetApdIdcStgwMasterResponse) {
	response = &SetApdIdcStgwMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置IDC绑定的安全网关Master
func (c *Client) SetApdIdcStgwMaster(request *SetApdIdcStgwMasterRequest) (response *SetApdIdcStgwMasterResponse, err error) {
	if request == nil {
		request = NewSetApdIdcStgwMasterRequest()
	}
	response = NewSetApdIdcStgwMasterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdApplyRequest() (request *DescribeApdApplyRequest) {
	request = &DescribeApdApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdApply")
	return
}

func NewDescribeApdApplyResponse() (response *DescribeApdApplyResponse) {
	response = &DescribeApdApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规则申请列表
func (c *Client) DescribeApdApply(request *DescribeApdApplyRequest) (response *DescribeApdApplyResponse, err error) {
	if request == nil {
		request = NewDescribeApdApplyRequest()
	}
	response = NewDescribeApdApplyResponse()
	err = c.Send(request, response)
	return
}

func NewSetApdLdStatusRequest() (request *SetApdLdStatusRequest) {
	request = &SetApdLdStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetApdLdStatus")
	return
}

func NewSetApdLdStatusResponse() (response *SetApdLdStatusResponse) {
	response = &SetApdLdStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用、隔离LD
func (c *Client) SetApdLdStatus(request *SetApdLdStatusRequest) (response *SetApdLdStatusResponse, err error) {
	if request == nil {
		request = NewSetApdLdStatusRequest()
	}
	response = NewSetApdLdStatusResponse()
	err = c.Send(request, response)
	return
}

func NewBindSetL4SetAndL7Request() (request *BindSetL4SetAndL7Request) {
	request = &BindSetL4SetAndL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "BindSetL4SetAndL7")
	return
}

func NewBindSetL4SetAndL7Response() (response *BindSetL4SetAndL7Response) {
	response = &BindSetL4SetAndL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 四层Set集群绑定7层Set集群
func (c *Client) BindSetL4SetAndL7(request *BindSetL4SetAndL7Request) (response *BindSetL4SetAndL7Response, err error) {
	if request == nil {
		request = NewBindSetL4SetAndL7Request()
	}
	response = NewBindSetL4SetAndL7Response()
	err = c.Send(request, response)
	return
}

func NewDescribeApdVipRequest() (request *DescribeApdVipRequest) {
	request = &DescribeApdVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdVip")
	return
}

func NewDescribeApdVipResponse() (response *DescribeApdVipResponse) {
	response = &DescribeApdVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询VIP详情
func (c *Client) DescribeApdVip(request *DescribeApdVipRequest) (response *DescribeApdVipResponse, err error) {
	if request == nil {
		request = NewDescribeApdVipRequest()
	}
	response = NewDescribeApdVipResponse()
	err = c.Send(request, response)
	return
}

func NewOpBgRsRequest() (request *OpBgRsRequest) {
	request = &OpBgRsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "OpBgRs")
	return
}

func NewOpBgRsResponse() (response *OpBgRsResponse) {
	response = &OpBgRsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置RS相关参数
func (c *Client) OpBgRs(request *OpBgRsRequest) (response *OpBgRsResponse, err error) {
	if request == nil {
		request = NewOpBgRsRequest()
	}
	response = NewOpBgRsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgCertificateRequest() (request *DescribeBgCertificateRequest) {
	request = &DescribeBgCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgCertificate")
	return
}

func NewDescribeBgCertificateResponse() (response *DescribeBgCertificateResponse) {
	response = &DescribeBgCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营侧证书详情
func (c *Client) DescribeBgCertificate(request *DescribeBgCertificateRequest) (response *DescribeBgCertificateResponse, err error) {
	if request == nil {
		request = NewDescribeBgCertificateRequest()
	}
	response = NewDescribeBgCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgServersRequest() (request *DescribeBgServersRequest) {
	request = &DescribeBgServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgServers")
	return
}

func NewDescribeBgServersResponse() (response *DescribeBgServersResponse) {
	response = &DescribeBgServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端服务器详情
func (c *Client) DescribeBgServers(request *DescribeBgServersRequest) (response *DescribeBgServersResponse, err error) {
	if request == nil {
		request = NewDescribeBgServersRequest()
	}
	response = NewDescribeBgServersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApdLdFromSetRequest() (request *DeleteApdLdFromSetRequest) {
	request = &DeleteApdLdFromSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteApdLdFromSet")
	return
}

func NewDeleteApdLdFromSetResponse() (response *DeleteApdLdFromSetResponse) {
	response = &DeleteApdLdFromSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Set集群中的LD
func (c *Client) DeleteApdLdFromSet(request *DeleteApdLdFromSetRequest) (response *DeleteApdLdFromSetResponse, err error) {
	if request == nil {
		request = NewDeleteApdLdFromSetRequest()
	}
	response = NewDeleteApdLdFromSetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApdVipRequest() (request *DeleteApdVipRequest) {
	request = &DeleteApdVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteApdVip")
	return
}

func NewDeleteApdVipResponse() (response *DeleteApdVipResponse) {
	response = &DeleteApdVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除VIP信息
func (c *Client) DeleteApdVip(request *DeleteApdVipRequest) (response *DeleteApdVipResponse, err error) {
	if request == nil {
		request = NewDeleteApdVipRequest()
	}
	response = NewDeleteApdVipResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBgRuleL7Request() (request *DeleteBgRuleL7Request) {
	request = &DeleteBgRuleL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteBgRuleL7")
	return
}

func NewDeleteBgRuleL7Response() (response *DeleteBgRuleL7Response) {
	response = &DeleteBgRuleL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除运营端7层规则
func (c *Client) DeleteBgRuleL7(request *DeleteBgRuleL7Request) (response *DeleteBgRuleL7Response, err error) {
	if request == nil {
		request = NewDeleteBgRuleL7Request()
	}
	response = NewDeleteBgRuleL7Response()
	err = c.Send(request, response)
	return
}

func NewQueryProductStateInfoRequest() (request *QueryProductStateInfoRequest) {
	request = &QueryProductStateInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "QueryProductStateInfo")
	return
}

func NewQueryProductStateInfoResponse() (response *QueryProductStateInfoResponse) {
	response = &QueryProductStateInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CLB产品状态信息查询接口
func (c *Client) QueryProductStateInfo(request *QueryProductStateInfoRequest) (response *QueryProductStateInfoResponse, err error) {
	if request == nil {
		request = NewQueryProductStateInfoRequest()
	}
	response = NewQueryProductStateInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBgRuleRequest() (request *DeleteBgRuleRequest) {
	request = &DeleteBgRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteBgRule")
	return
}

func NewDeleteBgRuleResponse() (response *DeleteBgRuleResponse) {
	response = &DeleteBgRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除bg规则
func (c *Client) DeleteBgRule(request *DeleteBgRuleRequest) (response *DeleteBgRuleResponse, err error) {
	if request == nil {
		request = NewDeleteBgRuleRequest()
	}
	response = NewDeleteBgRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgRuleRequest() (request *DescribeBgRuleRequest) {
	request = &DescribeBgRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgRule")
	return
}

func NewDescribeBgRuleResponse() (response *DescribeBgRuleResponse) {
	response = &DescribeBgRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端4层规则
func (c *Client) DescribeBgRule(request *DescribeBgRuleRequest) (response *DescribeBgRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBgRuleRequest()
	}
	response = NewDescribeBgRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSetLabelRequest() (request *DescribeSetLabelRequest) {
	request = &DescribeSetLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeSetLabel")
	return
}

func NewDescribeSetLabelResponse() (response *DescribeSetLabelResponse) {
	response = &DescribeSetLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据条件，获取集群和标签的关系
func (c *Client) DescribeSetLabel(request *DescribeSetLabelRequest) (response *DescribeSetLabelResponse, err error) {
	if request == nil {
		request = NewDescribeSetLabelRequest()
	}
	response = NewDescribeSetLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgTargetHealthRequest() (request *DescribeBgTargetHealthRequest) {
	request = &DescribeBgTargetHealthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgTargetHealth")
	return
}

func NewDescribeBgTargetHealthResponse() (response *DescribeBgTargetHealthResponse) {
	response = &DescribeBgTargetHealthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户clb实例健康状态
func (c *Client) DescribeBgTargetHealth(request *DescribeBgTargetHealthRequest) (response *DescribeBgTargetHealthResponse, err error) {
	if request == nil {
		request = NewDescribeBgTargetHealthRequest()
	}
	response = NewDescribeBgTargetHealthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdRuleRequest() (request *DescribeApdRuleRequest) {
	request = &DescribeApdRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdRule")
	return
}

func NewDescribeApdRuleResponse() (response *DescribeApdRuleResponse) {
	response = &DescribeApdRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询四层规则
func (c *Client) DescribeApdRule(request *DescribeApdRuleRequest) (response *DescribeApdRuleResponse, err error) {
	if request == nil {
		request = NewDescribeApdRuleRequest()
	}
	response = NewDescribeApdRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgLocationL7Request() (request *DescribeBgLocationL7Request) {
	request = &DescribeBgLocationL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgLocationL7")
	return
}

func NewDescribeBgLocationL7Response() (response *DescribeBgLocationL7Response) {
	response = &DescribeBgLocationL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询本地七层规则详情
func (c *Client) DescribeBgLocationL7(request *DescribeBgLocationL7Request) (response *DescribeBgLocationL7Response, err error) {
	if request == nil {
		request = NewDescribeBgLocationL7Request()
	}
	response = NewDescribeBgLocationL7Response()
	err = c.Send(request, response)
	return
}

func NewDescribeLdLocalIpRequest() (request *DescribeLdLocalIpRequest) {
	request = &DescribeLdLocalIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLdLocalIp")
	return
}

func NewDescribeLdLocalIpResponse() (response *DescribeLdLocalIpResponse) {
	response = &DescribeLdLocalIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询LD的LocalIP地址
func (c *Client) DescribeLdLocalIp(request *DescribeLdLocalIpRequest) (response *DescribeLdLocalIpResponse, err error) {
	if request == nil {
		request = NewDescribeLdLocalIpRequest()
	}
	response = NewDescribeLdLocalIpResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLbLimitRequest() (request *DescribeLbLimitRequest) {
	request = &DescribeLbLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLbLimit")
	return
}

func NewDescribeLbLimitResponse() (response *DescribeLbLimitResponse) {
	response = &DescribeLbLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询LB配置
func (c *Client) DescribeLbLimit(request *DescribeLbLimitRequest) (response *DescribeLbLimitResponse, err error) {
	if request == nil {
		request = NewDescribeLbLimitRequest()
	}
	response = NewDescribeLbLimitResponse()
	err = c.Send(request, response)
	return
}

func NewInitIspListRequest() (request *InitIspListRequest) {
	request = &InitIspListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "InitIspList")
	return
}

func NewInitIspListResponse() (response *InitIspListResponse) {
	response = &InitIspListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置运营商信息
func (c *Client) InitIspList(request *InitIspListRequest) (response *InitIspListResponse, err error) {
	if request == nil {
		request = NewInitIspListRequest()
	}
	response = NewInitIspListResponse()
	err = c.Send(request, response)
	return
}

func NewOpApdIdcRequest() (request *OpApdIdcRequest) {
	request = &OpApdIdcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "OpApdIdc")
	return
}

func NewOpApdIdcResponse() (response *OpApdIdcResponse) {
	response = &OpApdIdcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置物理机房参数
func (c *Client) OpApdIdc(request *OpApdIdcRequest) (response *OpApdIdcResponse, err error) {
	if request == nil {
		request = NewOpApdIdcRequest()
	}
	response = NewOpApdIdcResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApdVipGroupRequest() (request *UpdateApdVipGroupRequest) {
	request = &UpdateApdVipGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "UpdateApdVipGroup")
	return
}

func NewUpdateApdVipGroupResponse() (response *UpdateApdVipGroupResponse) {
	response = &UpdateApdVipGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新VIP组相关属性
func (c *Client) UpdateApdVipGroup(request *UpdateApdVipGroupRequest) (response *UpdateApdVipGroupResponse, err error) {
	if request == nil {
		request = NewUpdateApdVipGroupRequest()
	}
	response = NewUpdateApdVipGroupResponse()
	err = c.Send(request, response)
	return
}

func NewApplyBgRuleRequest() (request *ApplyBgRuleRequest) {
	request = &ApplyBgRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ApplyBgRule")
	return
}

func NewApplyBgRuleResponse() (response *ApplyBgRuleResponse) {
	response = &ApplyBgRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端规则申请
func (c *Client) ApplyBgRule(request *ApplyBgRuleRequest) (response *ApplyBgRuleResponse, err error) {
	if request == nil {
		request = NewApplyBgRuleRequest()
	}
	response = NewApplyBgRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApdLdRequest() (request *CreateApdLdRequest) {
	request = &CreateApdLdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateApdLd")
	return
}

func NewCreateApdLdResponse() (response *CreateApdLdResponse) {
	response = &CreateApdLdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用于部署负载均衡的服务器
func (c *Client) CreateApdLd(request *CreateApdLdRequest) (response *CreateApdLdResponse, err error) {
	if request == nil {
		request = NewCreateApdLdRequest()
	}
	response = NewCreateApdLdResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLdLocalIpRequest() (request *CreateLdLocalIpRequest) {
	request = &CreateLdLocalIpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateLdLocalIp")
	return
}

func NewCreateLdLocalIpResponse() (response *CreateLdLocalIpResponse) {
	response = &CreateLdLocalIpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// NAT64 Set集群绑定本地IP
func (c *Client) CreateLdLocalIp(request *CreateLdLocalIpRequest) (response *CreateLdLocalIpResponse, err error) {
	if request == nil {
		request = NewCreateLdLocalIpRequest()
	}
	response = NewCreateLdLocalIpResponse()
	err = c.Send(request, response)
	return
}

func NewSetApdIdcMasterRequest() (request *SetApdIdcMasterRequest) {
	request = &SetApdIdcMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetApdIdcMaster")
	return
}

func NewSetApdIdcMasterResponse() (response *SetApdIdcMasterResponse) {
	response = &SetApdIdcMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置IDC绑定的Master
func (c *Client) SetApdIdcMaster(request *SetApdIdcMasterRequest) (response *SetApdIdcMasterResponse, err error) {
	if request == nil {
		request = NewSetApdIdcMasterRequest()
	}
	response = NewSetApdIdcMasterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductFailureMigrateRequest() (request *CreateProductFailureMigrateRequest) {
	request = &CreateProductFailureMigrateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateProductFailureMigrate")
	return
}

func NewCreateProductFailureMigrateResponse() (response *CreateProductFailureMigrateResponse) {
	response = &CreateProductFailureMigrateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 产品故障切换操作接口
func (c *Client) CreateProductFailureMigrate(request *CreateProductFailureMigrateRequest) (response *CreateProductFailureMigrateResponse, err error) {
	if request == nil {
		request = NewCreateProductFailureMigrateRequest()
	}
	response = NewCreateProductFailureMigrateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdRuleL7Request() (request *DescribeApdRuleL7Request) {
	request = &DescribeApdRuleL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdRuleL7")
	return
}

func NewDescribeApdRuleL7Response() (response *DescribeApdRuleL7Response) {
	response = &DescribeApdRuleL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询7层规则详情
func (c *Client) DescribeApdRuleL7(request *DescribeApdRuleL7Request) (response *DescribeApdRuleL7Response, err error) {
	if request == nil {
		request = NewDescribeApdRuleL7Request()
	}
	response = NewDescribeApdRuleL7Response()
	err = c.Send(request, response)
	return
}

func NewOpBgL7RsRequest() (request *OpBgL7RsRequest) {
	request = &OpBgL7RsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "OpBgL7Rs")
	return
}

func NewOpBgL7RsResponse() (response *OpBgL7RsResponse) {
	response = &OpBgL7RsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 7层规则的RS 操作
func (c *Client) OpBgL7Rs(request *OpBgL7RsRequest) (response *OpBgL7RsResponse, err error) {
	if request == nil {
		request = NewOpBgL7RsRequest()
	}
	response = NewOpBgL7RsResponse()
	err = c.Send(request, response)
	return
}

func NewAssignApdLdToStgwSetRequest() (request *AssignApdLdToStgwSetRequest) {
	request = &AssignApdLdToStgwSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "AssignApdLdToStgwSet")
	return
}

func NewAssignApdLdToStgwSetResponse() (response *AssignApdLdToStgwSetResponse) {
	response = &AssignApdLdToStgwSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定LD到安全Set集群
func (c *Client) AssignApdLdToStgwSet(request *AssignApdLdToStgwSetRequest) (response *AssignApdLdToStgwSetResponse, err error) {
	if request == nil {
		request = NewAssignApdLdToStgwSetRequest()
	}
	response = NewAssignApdLdToStgwSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgJobRequest() (request *DescribeBgJobRequest) {
	request = &DescribeBgJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgJob")
	return
}

func NewDescribeBgJobResponse() (response *DescribeBgJobResponse) {
	response = &DescribeBgJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端任务详情
func (c *Client) DescribeBgJob(request *DescribeBgJobRequest) (response *DescribeBgJobResponse, err error) {
	if request == nil {
		request = NewDescribeBgJobRequest()
	}
	response = NewDescribeBgJobResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdIdcSetRequest() (request *DescribeApdIdcSetRequest) {
	request = &DescribeApdIdcSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdIdcSet")
	return
}

func NewDescribeApdIdcSetResponse() (response *DescribeApdIdcSetResponse) {
	response = &DescribeApdIdcSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询IDC Set集群属性
func (c *Client) DescribeApdIdcSet(request *DescribeApdIdcSetRequest) (response *DescribeApdIdcSetResponse, err error) {
	if request == nil {
		request = NewDescribeApdIdcSetRequest()
	}
	response = NewDescribeApdIdcSetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApdLdL7Request() (request *CreateApdLdL7Request) {
	request = &CreateApdLdL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateApdLdL7")
	return
}

func NewCreateApdLdL7Response() (response *CreateApdLdL7Response) {
	response = &CreateApdLdL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用于部署7层负载均衡的服务器
func (c *Client) CreateApdLdL7(request *CreateApdLdL7Request) (response *CreateApdLdL7Response, err error) {
	if request == nil {
		request = NewCreateApdLdL7Request()
	}
	response = NewCreateApdLdL7Response()
	err = c.Send(request, response)
	return
}

func NewCreateApdVipRequest() (request *CreateApdVipRequest) {
	request = &CreateApdVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateApdVip")
	return
}

func NewCreateApdVipResponse() (response *CreateApdVipResponse) {
	response = &CreateApdVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加VIP
func (c *Client) CreateApdVip(request *CreateApdVipRequest) (response *CreateApdVipResponse, err error) {
	if request == nil {
		request = NewCreateApdVipRequest()
	}
	response = NewCreateApdVipResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApdMasterRequest() (request *DeleteApdMasterRequest) {
	request = &DeleteApdMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteApdMaster")
	return
}

func NewDeleteApdMasterResponse() (response *DeleteApdMasterResponse) {
	response = &DeleteApdMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Master
func (c *Client) DeleteApdMaster(request *DeleteApdMasterRequest) (response *DeleteApdMasterResponse, err error) {
	if request == nil {
		request = NewDeleteApdMasterRequest()
	}
	response = NewDeleteApdMasterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdLdRequest() (request *DescribeApdLdRequest) {
	request = &DescribeApdLdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdLd")
	return
}

func NewDescribeApdLdResponse() (response *DescribeApdLdResponse) {
	response = &DescribeApdLdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Ld详细
func (c *Client) DescribeApdLd(request *DescribeApdLdRequest) (response *DescribeApdLdResponse, err error) {
	if request == nil {
		request = NewDescribeApdLdRequest()
	}
	response = NewDescribeApdLdResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateBgCertificateRequest() (request *UpdateBgCertificateRequest) {
	request = &UpdateBgCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "UpdateBgCertificate")
	return
}

func NewUpdateBgCertificateResponse() (response *UpdateBgCertificateResponse) {
	response = &UpdateBgCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新运营端证书相关属性
func (c *Client) UpdateBgCertificate(request *UpdateBgCertificateRequest) (response *UpdateBgCertificateResponse, err error) {
	if request == nil {
		request = NewUpdateBgCertificateRequest()
	}
	response = NewUpdateBgCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgRsByRuleRequest() (request *DescribeBgRsByRuleRequest) {
	request = &DescribeBgRsByRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgRsByRule")
	return
}

func NewDescribeBgRsByRuleResponse() (response *DescribeBgRsByRuleResponse) {
	response = &DescribeBgRsByRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 依据规则查询RS
func (c *Client) DescribeBgRsByRule(request *DescribeBgRsByRuleRequest) (response *DescribeBgRsByRuleResponse, err error) {
	if request == nil {
		request = NewDescribeBgRsByRuleRequest()
	}
	response = NewDescribeBgRsByRuleResponse()
	err = c.Send(request, response)
	return
}

func NewOpApdModuleRequest() (request *OpApdModuleRequest) {
	request = &OpApdModuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "OpApdModule")
	return
}

func NewOpApdModuleResponse() (response *OpApdModuleResponse) {
	response = &OpApdModuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置Module相关参数
func (c *Client) OpApdModule(request *OpApdModuleRequest) (response *OpApdModuleResponse, err error) {
	if request == nil {
		request = NewOpApdModuleRequest()
	}
	response = NewOpApdModuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApdLdFromStgwSetRequest() (request *DeleteApdLdFromStgwSetRequest) {
	request = &DeleteApdLdFromStgwSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteApdLdFromStgwSet")
	return
}

func NewDeleteApdLdFromStgwSetResponse() (response *DeleteApdLdFromStgwSetResponse) {
	response = &DeleteApdLdFromStgwSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除安全网关中的LD
func (c *Client) DeleteApdLdFromStgwSet(request *DeleteApdLdFromStgwSetRequest) (response *DeleteApdLdFromStgwSetResponse, err error) {
	if request == nil {
		request = NewDeleteApdLdFromStgwSetRequest()
	}
	response = NewDeleteApdLdFromStgwSetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdAppNameRequest() (request *DescribeApdAppNameRequest) {
	request = &DescribeApdAppNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdAppName")
	return
}

func NewDescribeApdAppNameResponse() (response *DescribeApdAppNameResponse) {
	response = &DescribeApdAppNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询申请业务名称
func (c *Client) DescribeApdAppName(request *DescribeApdAppNameRequest) (response *DescribeApdAppNameResponse, err error) {
	if request == nil {
		request = NewDescribeApdAppNameRequest()
	}
	response = NewDescribeApdAppNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgLoadBalancersRequest() (request *DescribeBgLoadBalancersRequest) {
	request = &DescribeBgLoadBalancersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgLoadBalancers")
	return
}

func NewDescribeBgLoadBalancersResponse() (response *DescribeBgLoadBalancersResponse) {
	response = &DescribeBgLoadBalancersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户资源信息
func (c *Client) DescribeBgLoadBalancers(request *DescribeBgLoadBalancersRequest) (response *DescribeBgLoadBalancersResponse, err error) {
	if request == nil {
		request = NewDescribeBgLoadBalancersRequest()
	}
	response = NewDescribeBgLoadBalancersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApdMasterRequest() (request *CreateApdMasterRequest) {
	request = &CreateApdMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateApdMaster")
	return
}

func NewCreateApdMasterResponse() (response *CreateApdMasterResponse) {
	response = &CreateApdMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Master环境
func (c *Client) CreateApdMaster(request *CreateApdMasterRequest) (response *CreateApdMasterResponse, err error) {
	if request == nil {
		request = NewCreateApdMasterRequest()
	}
	response = NewCreateApdMasterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApdVipGroupRequest() (request *CreateApdVipGroupRequest) {
	request = &CreateApdVipGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateApdVipGroup")
	return
}

func NewCreateApdVipGroupResponse() (response *CreateApdVipGroupResponse) {
	response = &CreateApdVipGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建VIP组
func (c *Client) CreateApdVipGroup(request *CreateApdVipGroupRequest) (response *CreateApdVipGroupResponse, err error) {
	if request == nil {
		request = NewCreateApdVipGroupRequest()
	}
	response = NewCreateApdVipGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgRsL7Request() (request *DescribeBgRsL7Request) {
	request = &DescribeBgRsL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgRsL7")
	return
}

func NewDescribeBgRsL7Response() (response *DescribeBgRsL7Response) {
	response = &DescribeBgRsL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询7层规则RS信息
func (c *Client) DescribeBgRsL7(request *DescribeBgRsL7Request) (response *DescribeBgRsL7Response, err error) {
	if request == nil {
		request = NewDescribeBgRsL7Request()
	}
	response = NewDescribeBgRsL7Response()
	err = c.Send(request, response)
	return
}

func NewDescribeL4RuleExRequest() (request *DescribeL4RuleExRequest) {
	request = &DescribeL4RuleExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeL4RuleEx")
	return
}

func NewDescribeL4RuleExResponse() (response *DescribeL4RuleExResponse) {
	response = &DescribeL4RuleExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询4层规则列表
func (c *Client) DescribeL4RuleEx(request *DescribeL4RuleExRequest) (response *DescribeL4RuleExResponse, err error) {
	if request == nil {
		request = NewDescribeL4RuleExRequest()
	}
	response = NewDescribeL4RuleExResponse()
	err = c.Send(request, response)
	return
}

func NewExportApdRuleRequest() (request *ExportApdRuleRequest) {
	request = &ExportApdRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ExportApdRule")
	return
}

func NewExportApdRuleResponse() (response *ExportApdRuleResponse) {
	response = &ExportApdRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出四层规则
func (c *Client) ExportApdRule(request *ExportApdRuleRequest) (response *ExportApdRuleResponse, err error) {
	if request == nil {
		request = NewExportApdRuleRequest()
	}
	response = NewExportApdRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllIspListRequest() (request *DescribeAllIspListRequest) {
	request = &DescribeAllIspListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeAllIspList")
	return
}

func NewDescribeAllIspListResponse() (response *DescribeAllIspListResponse) {
	response = &DescribeAllIspListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有运营商信息
func (c *Client) DescribeAllIspList(request *DescribeAllIspListRequest) (response *DescribeAllIspListResponse, err error) {
	if request == nil {
		request = NewDescribeAllIspListRequest()
	}
	response = NewDescribeAllIspListResponse()
	err = c.Send(request, response)
	return
}

func NewOpApdSetRequest() (request *OpApdSetRequest) {
	request = &OpApdSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "OpApdSet")
	return
}

func NewOpApdSetResponse() (response *OpApdSetResponse) {
	response = &OpApdSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置set集群相关参数
func (c *Client) OpApdSet(request *OpApdSetRequest) (response *OpApdSetResponse, err error) {
	if request == nil {
		request = NewOpApdSetRequest()
	}
	response = NewOpApdSetResponse()
	err = c.Send(request, response)
	return
}

func NewBindAppIdLabelRequest() (request *BindAppIdLabelRequest) {
	request = &BindAppIdLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "BindAppIdLabel")
	return
}

func NewBindAppIdLabelResponse() (response *BindAppIdLabelResponse) {
	response = &BindAppIdLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定用户和标签
func (c *Client) BindAppIdLabel(request *BindAppIdLabelRequest) (response *BindAppIdLabelResponse, err error) {
	if request == nil {
		request = NewBindAppIdLabelRequest()
	}
	response = NewBindAppIdLabelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApdMasterRequest() (request *ModifyApdMasterRequest) {
	request = &ModifyApdMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyApdMaster")
	return
}

func NewModifyApdMasterResponse() (response *ModifyApdMasterResponse) {
	response = &ModifyApdMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Master环境配置
func (c *Client) ModifyApdMaster(request *ModifyApdMasterRequest) (response *ModifyApdMasterResponse, err error) {
	if request == nil {
		request = NewModifyApdMasterRequest()
	}
	response = NewModifyApdMasterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdMasterRequest() (request *DescribeApdMasterRequest) {
	request = &DescribeApdMasterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdMaster")
	return
}

func NewDescribeApdMasterResponse() (response *DescribeApdMasterResponse) {
	response = &DescribeApdMasterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Master详情
func (c *Client) DescribeApdMaster(request *DescribeApdMasterRequest) (response *DescribeApdMasterResponse, err error) {
	if request == nil {
		request = NewDescribeApdMasterRequest()
	}
	response = NewDescribeApdMasterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLbLimitRequest() (request *ModifyLbLimitRequest) {
	request = &ModifyLbLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyLbLimit")
	return
}

func NewModifyLbLimitResponse() (response *ModifyLbLimitResponse) {
	response = &ModifyLbLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改LB配额限制。
func (c *Client) ModifyLbLimit(request *ModifyLbLimitRequest) (response *ModifyLbLimitResponse, err error) {
	if request == nil {
		request = NewModifyLbLimitRequest()
	}
	response = NewModifyLbLimitResponse()
	err = c.Send(request, response)
	return
}

func NewSetBgRuleRequest() (request *SetBgRuleRequest) {
	request = &SetBgRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetBgRule")
	return
}

func NewSetBgRuleResponse() (response *SetBgRuleResponse) {
	response = &SetBgRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置四层规则
func (c *Client) SetBgRule(request *SetBgRuleRequest) (response *SetBgRuleResponse, err error) {
	if request == nil {
		request = NewSetBgRuleRequest()
	}
	response = NewSetBgRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDefaultIspListRequest() (request *UpdateDefaultIspListRequest) {
	request = &UpdateDefaultIspListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "UpdateDefaultIspList")
	return
}

func NewUpdateDefaultIspListResponse() (response *UpdateDefaultIspListResponse) {
	response = &UpdateDefaultIspListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新默认运营商
func (c *Client) UpdateDefaultIspList(request *UpdateDefaultIspListRequest) (response *UpdateDefaultIspListResponse, err error) {
	if request == nil {
		request = NewUpdateDefaultIspListRequest()
	}
	response = NewUpdateDefaultIspListResponse()
	err = c.Send(request, response)
	return
}

func NewExportApdRulePubCertRequest() (request *ExportApdRulePubCertRequest) {
	request = &ExportApdRulePubCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ExportApdRulePubCert")
	return
}

func NewExportApdRulePubCertResponse() (response *ExportApdRulePubCertResponse) {
	response = &ExportApdRulePubCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出规则证书
func (c *Client) ExportApdRulePubCert(request *ExportApdRulePubCertRequest) (response *ExportApdRulePubCertResponse, err error) {
	if request == nil {
		request = NewExportApdRulePubCertRequest()
	}
	response = NewExportApdRulePubCertResponse()
	err = c.Send(request, response)
	return
}

func NewUploadBgCertificateRequest() (request *UploadBgCertificateRequest) {
	request = &UploadBgCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "UploadBgCertificate")
	return
}

func NewUploadBgCertificateResponse() (response *UploadBgCertificateResponse) {
	response = &UploadBgCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上传运营端证书
func (c *Client) UploadBgCertificate(request *UploadBgCertificateRequest) (response *UploadBgCertificateResponse, err error) {
	if request == nil {
		request = NewUploadBgCertificateRequest()
	}
	response = NewUploadBgCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdCertificateRequest() (request *DescribeApdCertificateRequest) {
	request = &DescribeApdCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdCertificate")
	return
}

func NewDescribeApdCertificateResponse() (response *DescribeApdCertificateResponse) {
	response = &DescribeApdCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询证书详情
func (c *Client) DescribeApdCertificate(request *DescribeApdCertificateRequest) (response *DescribeApdCertificateResponse, err error) {
	if request == nil {
		request = NewDescribeApdCertificateRequest()
	}
	response = NewDescribeApdCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApdLdRequest() (request *DeleteApdLdRequest) {
	request = &DeleteApdLdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteApdLd")
	return
}

func NewDeleteApdLdResponse() (response *DeleteApdLdResponse) {
	response = &DeleteApdLdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除LD
func (c *Client) DeleteApdLd(request *DeleteApdLdRequest) (response *DeleteApdLdResponse, err error) {
	if request == nil {
		request = NewDeleteApdLdRequest()
	}
	response = NewDeleteApdLdResponse()
	err = c.Send(request, response)
	return
}

func NewExportApdRuleL7Request() (request *ExportApdRuleL7Request) {
	request = &ExportApdRuleL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ExportApdRuleL7")
	return
}

func NewExportApdRuleL7Response() (response *ExportApdRuleL7Response) {
	response = &ExportApdRuleL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出七层规则
func (c *Client) ExportApdRuleL7(request *ExportApdRuleL7Request) (response *ExportApdRuleL7Response, err error) {
	if request == nil {
		request = NewExportApdRuleL7Request()
	}
	response = NewExportApdRuleL7Response()
	err = c.Send(request, response)
	return
}

func NewDeleteBgCertificateRequest() (request *DeleteBgCertificateRequest) {
	request = &DeleteBgCertificateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteBgCertificate")
	return
}

func NewDeleteBgCertificateResponse() (response *DeleteBgCertificateResponse) {
	response = &DeleteBgCertificateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除证书信息
func (c *Client) DeleteBgCertificate(request *DeleteBgCertificateRequest) (response *DeleteBgCertificateResponse, err error) {
	if request == nil {
		request = NewDeleteBgCertificateRequest()
	}
	response = NewDeleteBgCertificateResponse()
	err = c.Send(request, response)
	return
}

func NewExportBgCertListRequest() (request *ExportBgCertListRequest) {
	request = &ExportBgCertListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ExportBgCertList")
	return
}

func NewExportBgCertListResponse() (response *ExportBgCertListResponse) {
	response = &ExportBgCertListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出证书列表
func (c *Client) ExportBgCertList(request *ExportBgCertListRequest) (response *ExportBgCertListResponse, err error) {
	if request == nil {
		request = NewExportBgCertListRequest()
	}
	response = NewExportBgCertListResponse()
	err = c.Send(request, response)
	return
}

func NewSetBgL7RuleRequest() (request *SetBgL7RuleRequest) {
	request = &SetBgL7RuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetBgL7Rule")
	return
}

func NewSetBgL7RuleResponse() (response *SetBgL7RuleResponse) {
	response = &SetBgL7RuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置7层规则参数
func (c *Client) SetBgL7Rule(request *SetBgL7RuleRequest) (response *SetBgL7RuleResponse, err error) {
	if request == nil {
		request = NewSetBgL7RuleRequest()
	}
	response = NewSetBgL7RuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppIdLabelRequest() (request *DescribeAppIdLabelRequest) {
	request = &DescribeAppIdLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeAppIdLabel")
	return
}

func NewDescribeAppIdLabelResponse() (response *DescribeAppIdLabelResponse) {
	response = &DescribeAppIdLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询AppId和标签的关系
func (c *Client) DescribeAppIdLabel(request *DescribeAppIdLabelRequest) (response *DescribeAppIdLabelResponse, err error) {
	if request == nil {
		request = NewDescribeAppIdLabelRequest()
	}
	response = NewDescribeAppIdLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApdVipSegmentRequest() (request *DeleteApdVipSegmentRequest) {
	request = &DeleteApdVipSegmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteApdVipSegment")
	return
}

func NewDeleteApdVipSegmentResponse() (response *DeleteApdVipSegmentResponse) {
	response = &DeleteApdVipSegmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除VIP整个网段
func (c *Client) DeleteApdVipSegment(request *DeleteApdVipSegmentRequest) (response *DeleteApdVipSegmentResponse, err error) {
	if request == nil {
		request = NewDeleteApdVipSegmentRequest()
	}
	response = NewDeleteApdVipSegmentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeL7RuleExRequest() (request *DescribeL7RuleExRequest) {
	request = &DescribeL7RuleExRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeL7RuleEx")
	return
}

func NewDescribeL7RuleExResponse() (response *DescribeL7RuleExResponse) {
	response = &DescribeL7RuleExResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询7层规则扩展
func (c *Client) DescribeL7RuleEx(request *DescribeL7RuleExRequest) (response *DescribeL7RuleExResponse, err error) {
	if request == nil {
		request = NewDescribeL7RuleExRequest()
	}
	response = NewDescribeL7RuleExResponse()
	err = c.Send(request, response)
	return
}

func NewRetireBgApplyRequest() (request *RetireBgApplyRequest) {
	request = &RetireBgApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "RetireBgApply")
	return
}

func NewRetireBgApplyResponse() (response *RetireBgApplyResponse) {
	response = &RetireBgApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 驳回运营端规则申请
func (c *Client) RetireBgApply(request *RetireBgApplyRequest) (response *RetireBgApplyResponse, err error) {
	if request == nil {
		request = NewRetireBgApplyRequest()
	}
	response = NewRetireBgApplyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgRuleL7Request() (request *DescribeBgRuleL7Request) {
	request = &DescribeBgRuleL7Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgRuleL7")
	return
}

func NewDescribeBgRuleL7Response() (response *DescribeBgRuleL7Response) {
	response = &DescribeBgRuleL7Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运营端7层规则详情
func (c *Client) DescribeBgRuleL7(request *DescribeBgRuleL7Request) (response *DescribeBgRuleL7Response, err error) {
	if request == nil {
		request = NewDescribeBgRuleL7Request()
	}
	response = NewDescribeBgRuleL7Response()
	err = c.Send(request, response)
	return
}

func NewOpApdSectionRequest() (request *OpApdSectionRequest) {
	request = &OpApdSectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "OpApdSection")
	return
}

func NewOpApdSectionResponse() (response *OpApdSectionResponse) {
	response = &OpApdSectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 专区相关操作
func (c *Client) OpApdSection(request *OpApdSectionRequest) (response *OpApdSectionResponse, err error) {
	if request == nil {
		request = NewOpApdSectionRequest()
	}
	response = NewOpApdSectionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApdMasterListRequest() (request *DescribeApdMasterListRequest) {
	request = &DescribeApdMasterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeApdMasterList")
	return
}

func NewDescribeApdMasterListResponse() (response *DescribeApdMasterListResponse) {
	response = &DescribeApdMasterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Master列表详情
func (c *Client) DescribeApdMasterList(request *DescribeApdMasterListRequest) (response *DescribeApdMasterListResponse, err error) {
	if request == nil {
		request = NewDescribeApdMasterListRequest()
	}
	response = NewDescribeApdMasterListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBgApplyRequest() (request *DescribeBgApplyRequest) {
	request = &DescribeBgApplyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBgApply")
	return
}

func NewDescribeBgApplyResponse() (response *DescribeBgApplyResponse) {
	response = &DescribeBgApplyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询规则申请记录
func (c *Client) DescribeBgApply(request *DescribeBgApplyRequest) (response *DescribeBgApplyResponse, err error) {
	if request == nil {
		request = NewDescribeBgApplyRequest()
	}
	response = NewDescribeBgApplyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryProductFailureMigrateTaskDetailRequest() (request *QueryProductFailureMigrateTaskDetailRequest) {
	request = &QueryProductFailureMigrateTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "QueryProductFailureMigrateTaskDetail")
	return
}

func NewQueryProductFailureMigrateTaskDetailResponse() (response *QueryProductFailureMigrateTaskDetailResponse) {
	response = &QueryProductFailureMigrateTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口主要用于查询产品故障转移结果。
func (c *Client) QueryProductFailureMigrateTaskDetail(request *QueryProductFailureMigrateTaskDetailRequest) (response *QueryProductFailureMigrateTaskDetailResponse, err error) {
	if request == nil {
		request = NewQueryProductFailureMigrateTaskDetailRequest()
	}
	response = NewQueryProductFailureMigrateTaskDetailResponse()
	err = c.Send(request, response)
	return
}
