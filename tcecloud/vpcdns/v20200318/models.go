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

package v20200318

import (
	"encoding/json"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

// to suppress unused import error, although ugly
var (
	_ = tchttp.POST
	_ = json.Marshal
)

type ForwardRecord struct {
	// 记录ID

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Forward DNS

	ForwardDns *string `json:"ForwardDns,omitempty" name:"ForwardDns"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 最后更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DescribeForwardRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Forward DNS 记录值列表

		Records []*ForwardRecord `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeForwardRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeForwardRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsLogStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVpcDnsLogStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsLogStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListFilterItem struct {
	// 过滤条件: `Domain`解析域; `VpcInfo` ， VPC ID或者实例ID; Account, 用户OwnerUin

	Name *string `json:"Name,omitempty" name:"Name"`
	// 过滤条件值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeVpcDnsClusterListRequest struct {
	*tchttp.BaseRequest

	// 分页大小，默认20，最大2000

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量，从0开始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeVpcDnsClusterListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsClusterListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCountInfo struct {
	// 集群总数

	AllTotal *int64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 集群总数

	ClusterTotal *int64 `json:"ClusterTotal,omitempty" name:"ClusterTotal"`
}

type ClusterRegion struct {
	// 地区

	Region *string `json:"Region,omitempty" name:"Region"`
}

type RecordDetail struct {
	// 记录ID

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
	// 解析域ID

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 子域名

	SubDomain *string `json:"SubDomain,omitempty" name:"SubDomain"`
	// 记录类型， A, AAAA, CNAME, MX, PTR

	RecordType *string `json:"RecordType,omitempty" name:"RecordType"`
	// 记录值

	Value *bool `json:"Value,omitempty" name:"Value"`
	// 记录过期时间

	Ttl *int64 `json:"Ttl,omitempty" name:"Ttl"`
	// MX优先级

	Mx *int64 `json:"Mx,omitempty" name:"Mx"`
	// 记录是否启用

	Enabled *int64 `json:"Enabled,omitempty" name:"Enabled"`
	// 记录状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	// 记录权重

	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

type DescribeOcloudVpcDnsDomainListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表统计信息

		Info []*DomainCountInfo `json:"Info,omitempty" name:"Info"`
		// 解析域列表

		Domains []*DomainDetail `json:"Domains,omitempty" name:"Domains"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOcloudVpcDnsDomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOcloudVpcDnsDomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsLogStatusRequest struct {
	*tchttp.BaseRequest

	// 日志开关状态, On：打开，Off ：关闭

	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyVpcDnsLogStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsLogStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID数组

	ClusterIds []*int64 `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DeleteVpcDnsClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群成员

	ClusterMember []*string `json:"ClusterMember,omitempty" name:"ClusterMember"`
	// 集群所属区域

	ClusterRegion *ClusterRegion `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *ModifyVpcDnsClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVpcDnsClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateVpcDnsClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcdnsClusterInfo struct {
	// 集群ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 所属区域

	ClusterRegion []*ClusterRegion `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
	// 集群成员

	ClusterMember []*string `json:"ClusterMember,omitempty" name:"ClusterMember"`
	// 归属人UIN

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 操作人UIN

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
}

type ModifyForwardRecordRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// foward dns

	ForwardDns *string `json:"ForwardDns,omitempty" name:"ForwardDns"`
	// Record Id

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
}

func (r *ModifyForwardRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyForwardRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCountInfo struct {
	// 域名总数

	AllTotal *int64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 域名总数

	DomainTotal *int64 `json:"DomainTotal,omitempty" name:"DomainTotal"`
}

type DomainDetail struct {
	// 解析域ID

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 归属UIN

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 解析域域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 创建时间

	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`
	// 修改时间

	UpdatedOn *string `json:"UpdatedOn,omitempty" name:"UpdatedOn"`
	// 解析域记录数

	RecordCount *int64 `json:"RecordCount,omitempty" name:"RecordCount"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 绑定的VPC列表

	VpcInfos *VpcInfo `json:"VpcInfos,omitempty" name:"VpcInfos"`
}

type CreateForwardRecordRequest struct {
	*tchttp.BaseRequest

	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Foward DNS

	ForwardDns *string `json:"ForwardDns,omitempty" name:"ForwardDns"`
}

func (r *CreateForwardRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateForwardRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordCountInfo struct {
	// 记录数

	AllTotal *int64 `json:"AllTotal,omitempty" name:"AllTotal"`
	// 记录数

	RecordTotal *int64 `json:"RecordTotal,omitempty" name:"RecordTotal"`
}

type ModifyVpcDnsLogStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVpcDnsLogStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVpcDnsLogStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {
	// VPC实例ID

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
	// VPC ID

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// 区域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
}

type DeleteForwardRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteForwardRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteForwardRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOcloudVpcDnsDomainListRequest struct {
	*tchttp.BaseRequest

	// 分页大小，默认20，最大200

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量，从0开始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*ListFilterItem `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOcloudVpcDnsDomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOcloudVpcDnsDomainListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOcloudVpcDnsRecordListRequest struct {
	*tchttp.BaseRequest

	// 解析域id

	DomainId *int64 `json:"DomainId,omitempty" name:"DomainId"`
	// 分页大小，默认20，最大200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页偏移量，从0开始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件

	Filters []*ListFilterItem `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOcloudVpcDnsRecordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOcloudVpcDnsRecordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteVpcDnsClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteVpcDnsClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteVpcDnsClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOcloudVpcDnsRecordListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表统计信息

		Info []*RecordCountInfo `json:"Info,omitempty" name:"Info"`
		// 记录详情列表

		Records []*RecordDetail `json:"Records,omitempty" name:"Records"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOcloudVpcDnsRecordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOcloudVpcDnsRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsClusterListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表统计信息

		Info *ClusterCountInfo `json:"Info,omitempty" name:"Info"`
		// 集群列表

		Clusters *VpcdnsClusterInfo `json:"Clusters,omitempty" name:"Clusters"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsClusterListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyForwardRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyForwardRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyForwardRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeForwardRecordsRequest struct {
	*tchttp.BaseRequest

	// 分页偏移量，从0开始

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页大小，默认20，最大200

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// Forward DNS

	ForwardDns *string `json:"ForwardDns,omitempty" name:"ForwardDns"`
}

func (r *DescribeForwardRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeForwardRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateForwardRecordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateForwardRecordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateForwardRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateVpcDnsClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群成员IP列表

	ClusterMember []*string `json:"ClusterMember,omitempty" name:"ClusterMember"`
	// 集群所属区域

	ClusterRegion *ClusterRegion `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *CreateVpcDnsClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateVpcDnsClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteForwardRecordRequest struct {
	*tchttp.BaseRequest

	// 记录ID

	RecordId *int64 `json:"RecordId,omitempty" name:"RecordId"`
}

func (r *DeleteForwardRecordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteForwardRecordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVpcDnsLogStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 解析日志开启状态:  On, Off

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVpcDnsLogStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVpcDnsLogStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
