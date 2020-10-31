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

// Response Object
type ShowL7policyResponse struct {
	L7policy *L7policyV2Resp `json:"l7policy,omitempty"`
}

func (o ShowL7policyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowL7policyResponse", string(data)}, " ")
}