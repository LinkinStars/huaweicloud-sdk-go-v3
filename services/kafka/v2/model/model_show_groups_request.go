/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowGroupsRequest struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	Group      string `json:"group"`
}

func (o ShowGroupsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowGroupsRequest", string(data)}, " ")
}
