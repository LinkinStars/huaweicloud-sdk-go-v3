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
type UpdateUserRequestBody struct {
	User *UpdateUserOption `json:"user"`
}

func (o UpdateUserRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateUserRequestBody", string(data)}, " ")
}
