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

// Response Object
type ShowPartitionBeginningMessageResponse struct {
	// Topic名称。
	Topic *string `json:"topic,omitempty"`
	// 分区编号。
	Partition *int32 `json:"partition,omitempty"`
	// 最新消息位置。
	MessageOffset *int32 `json:"message_offset,omitempty"`
}

func (o ShowPartitionBeginningMessageResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPartitionBeginningMessageResponse", string(data)}, " ")
}
