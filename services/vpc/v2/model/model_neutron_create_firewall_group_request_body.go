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

//
type NeutronCreateFirewallGroupRequestBody struct {
	FirewallGroup *NeutronCreateFirewallGroupOption `json:"firewall_group"`
}

func (o NeutronCreateFirewallGroupRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronCreateFirewallGroupRequestBody", string(data)}, " ")
}
