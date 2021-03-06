/*
 * NAT
 *
 * Open Api of Public Nat.
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListNatGatewaySnatRulesResponse struct {
	// 查询SNAT规则列表的响应体。
	SnatRules *[]NatGatewaySnatRuleResponseBody `json:"snat_rules,omitempty"`
}

func (o ListNatGatewaySnatRulesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListNatGatewaySnatRulesResponse", string(data)}, " ")
}
