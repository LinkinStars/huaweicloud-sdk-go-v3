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
type ConfirmConsumptionMessagesRequest struct {
	ProjectId       string                         `json:"project_id"`
	QueueId         string                         `json:"queue_id"`
	ConsumerGroupId string                         `json:"consumer_group_id"`
	Body            *ConfirmConsumptionMessagesReq `json:"body,omitempty"`
}

func (o ConfirmConsumptionMessagesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ConfirmConsumptionMessagesRequest", string(data)}, " ")
}
