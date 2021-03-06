/*
 * RabbitMQ
 *
 * RabbitMQ Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowInstanceTagsResponse struct {
	// 标签列表
	Tags *[]CreateInstanceReqTags `json:"tags,omitempty"`
}

func (o ShowInstanceTagsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowInstanceTagsResponse", string(data)}, " ")
}
