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

// This is a auto create Body Object
type CreateLoadbalancerRequestBody struct {
	Loadbalancer *CreateLoadbalancerReq `json:"loadbalancer"`
}

func (o CreateLoadbalancerRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateLoadbalancerRequestBody", string(data)}, " ")
}
