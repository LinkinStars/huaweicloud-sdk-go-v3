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

// 创建SNAT规则的请求体。
type CreateNatGatewaySnatRuleRequestOption struct {
	SnatRule *CreateNatGatewaySnatRuleOption `json:"snat_rule"`
}

func (o CreateNatGatewaySnatRuleRequestOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateNatGatewaySnatRuleRequestOption", string(data)}, " ")
}
