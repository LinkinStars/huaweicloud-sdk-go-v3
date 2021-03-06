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
type NeutronCreateFirewallRuleResponse struct {
	FirewallRule *NeutronFirewallRule `json:"firewall_rule,omitempty"`
}

func (o NeutronCreateFirewallRuleResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronCreateFirewallRuleResponse", string(data)}, " ")
}
