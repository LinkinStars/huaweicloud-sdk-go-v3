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
type BatchCreateOrDeleteInstanceTagRequest struct {
	ProjectId  string                     `json:"project_id"`
	InstanceId string                     `json:"instance_id"`
	Body       *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteInstanceTagRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateOrDeleteInstanceTagRequest", string(data)}, " ")
}
