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

// Response Object
type CreateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
}

func (o CreateSecurityGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSecurityGroupResponse", string(data)}, " ")
}
