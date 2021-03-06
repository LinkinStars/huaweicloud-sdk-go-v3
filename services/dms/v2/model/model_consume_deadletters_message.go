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

type ConsumeDeadlettersMessage struct {
	Message *ConsumeDeadlettersMessageMessage `json:"message,omitempty"`
	// 消息handler。
	Handler *string `json:"handler,omitempty"`
}

func (o ConsumeDeadlettersMessage) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ConsumeDeadlettersMessage", string(data)}, " ")
}
