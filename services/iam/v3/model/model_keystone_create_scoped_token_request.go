/*
 * IAM
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
type KeystoneCreateScopedTokenRequest struct {
	Body *KeystoneCreateScopedTokenRequestBody `json:"body,omitempty"`
}

func (o KeystoneCreateScopedTokenRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateScopedTokenRequest", string(data)}, " ")
}
