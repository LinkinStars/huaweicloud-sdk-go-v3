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
type UpdateIpGroupRequest struct {
	IpgroupId string                    `json:"ipgroup_id"`
	Body      *UpdateIpGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateIpGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateIpGroupRequest", string(data)}, " ")
}
