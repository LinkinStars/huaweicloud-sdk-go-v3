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
type CountPreoccupyIpNumRequest struct {
	L7FlavorId             *string `json:"l7_flavor_id,omitempty"`
	IpTargetEnable         *bool   `json:"ip_target_enable,omitempty"`
	IpVersion              *bool   `json:"ip_version,omitempty"`
	LoadbalancerId         *string `json:"loadbalancer_id,omitempty"`
	AvailabilityZoneNumber *int32  `json:"availability_zone_number,omitempty"`
}

func (o CountPreoccupyIpNumRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CountPreoccupyIpNumRequest", string(data)}, " ")
}