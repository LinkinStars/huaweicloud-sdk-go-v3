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

type ConfirmDeadLettersMessagesReq struct {
	// 确认消息数组。
	Message *[]ConfirmDeadLettersMessagesReqMessage `json:"message,omitempty"`
}

func (o ConfirmDeadLettersMessagesReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ConfirmDeadLettersMessagesReq", string(data)}, " ")
}
