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

// 更新SNAT规则的请求体。
type UpdateNatGatewaySnatRuleRequestOption struct {
	SnatRule *UpdateNatGatewaySnatRuleOption `json:"snat_rule"`
}

func (o UpdateNatGatewaySnatRuleRequestOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateNatGatewaySnatRuleRequestOption", string(data)}, " ")
}
