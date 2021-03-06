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
type KeystoneCreateUserTokenByPasswordAndMfaRequestBody struct {
	Auth *MfaAuth `json:"auth"`
}

func (o KeystoneCreateUserTokenByPasswordAndMfaRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordAndMfaRequestBody", string(data)}, " ")
}
