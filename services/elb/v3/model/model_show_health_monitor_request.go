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
type ShowHealthMonitorRequest struct {
	HealthmonitorId string `json:"healthmonitor_id"`
}

func (o ShowHealthMonitorRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowHealthMonitorRequest", string(data)}, " ")
}
