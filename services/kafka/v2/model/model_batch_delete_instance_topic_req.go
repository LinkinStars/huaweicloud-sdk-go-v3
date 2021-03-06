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

type BatchDeleteInstanceTopicReq struct {
	// 待删除的topic列表。
	Topics *[]string `json:"topics,omitempty"`
}

func (o BatchDeleteInstanceTopicReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteInstanceTopicReq", string(data)}, " ")
}
