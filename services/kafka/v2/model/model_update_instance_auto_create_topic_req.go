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

type UpdateInstanceAutoCreateTopicReq struct {
	// 是否开启自动创建topic功能。
	EnableAutoTopic bool `json:"enable_auto_topic"`
}

func (o UpdateInstanceAutoCreateTopicReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateInstanceAutoCreateTopicReq", string(data)}, " ")
}
