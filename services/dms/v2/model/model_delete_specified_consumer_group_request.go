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
type DeleteSpecifiedConsumerGroupRequest struct {
	ProjectId string `json:"project_id"`
	QueueId   string `json:"queue_id"`
	GroupId   string `json:"group_id"`
}

func (o DeleteSpecifiedConsumerGroupRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteSpecifiedConsumerGroupRequest", string(data)}, " ")
}
