/*
 * DMS
 *
 * DMS Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SendMessagesRequest struct {
	ProjectId string           `json:"project_id"`
	QueueId   string           `json:"queue_id"`
	Body      *SendMessagesReq `json:"body,omitempty"`
}

func (o SendMessagesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SendMessagesRequest", string(data)}, " ")
}
