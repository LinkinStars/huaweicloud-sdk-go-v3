/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSubnetRequest struct {
	VpcId    string                   `json:"vpc_id"`
	SubnetId string                   `json:"subnet_id"`
	Body     *UpdateSubnetRequestBody `json:"body,omitempty"`
}

func (o UpdateSubnetRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateSubnetRequest", string(data)}, " ")
}
