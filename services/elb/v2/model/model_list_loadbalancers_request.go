/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListLoadbalancersRequest struct {
	Limit               *int32  `json:"limit,omitempty"`
	Marker              *string `json:"marker,omitempty"`
	PageReverse         *bool   `json:"page_reverse,omitempty"`
	Id                  *string `json:"id,omitempty"`
	Description         *string `json:"description,omitempty"`
	Name                *string `json:"name,omitempty"`
	OperatingStatus     *string `json:"operating_status,omitempty"`
	ProvisioningStatus  *string `json:"provisioning_status,omitempty"`
	VipAddress          *string `json:"vip_address,omitempty"`
	VipPortId           *string `json:"vip_port_id,omitempty"`
	VipSubnetId         *string `json:"vip_subnet_id,omitempty"`
	VpcId               *string `json:"vpc_id,omitempty"`
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	AdminStateUp        *bool   `json:"admin_state_up,omitempty"`
	MemberAddress       *string `json:"member_address,omitempty"`
	MemberDeviceId      *string `json:"member_device_id,omitempty"`
}

func (o ListLoadbalancersRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListLoadbalancersRequest", string(data)}, " ")
}
