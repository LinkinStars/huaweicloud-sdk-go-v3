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

//
type KeystoneUpdateProtocolRequestBody struct {
	Protocol *ProtocolOption `json:"protocol"`
}

func (o KeystoneUpdateProtocolRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateProtocolRequestBody", string(data)}, " ")
}
