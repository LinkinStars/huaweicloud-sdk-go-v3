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

// 创建DNAT规则的请求体。
type CreateNatGatewayDnatRuleOption struct {
	DnatRule *CreateNatGatewayDnatOption `json:"dnat_rule"`
}

func (o CreateNatGatewayDnatRuleOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNatGatewayDnatRuleOption", string(data)}, " ")
}
