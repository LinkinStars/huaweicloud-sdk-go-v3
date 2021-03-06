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
type ShowLoadbalancersStatusResponse struct {
	Statuses *StatusResp `json:"statuses,omitempty"`
}

func (o ShowLoadbalancersStatusResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowLoadbalancersStatusResponse", string(data)}, " ")
}
