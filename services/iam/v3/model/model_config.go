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
type Config struct {
	SecurityCompliance *SecurityCompliance `json:"security_compliance"`
}

func (o Config) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Config", string(data)}, " ")
}
