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

// Request Object
type CreateWhitelistRequest struct {
	Body *CreateWhitelistRequestBody `json:"body,omitempty"`
}

func (o CreateWhitelistRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateWhitelistRequest", string(data)}, " ")
}
