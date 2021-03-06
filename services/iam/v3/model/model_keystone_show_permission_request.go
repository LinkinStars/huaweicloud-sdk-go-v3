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
type KeystoneShowPermissionRequest struct {
	RoleId string `json:"role_id"`
}

func (o KeystoneShowPermissionRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneShowPermissionRequest", string(data)}, " ")
}
