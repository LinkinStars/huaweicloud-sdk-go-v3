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
type DeleteListenerResponse struct {
}

func (o DeleteListenerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteListenerResponse", string(data)}, " ")
}
