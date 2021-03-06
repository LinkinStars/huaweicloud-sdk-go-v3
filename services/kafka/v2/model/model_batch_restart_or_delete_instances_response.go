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
type BatchRestartOrDeleteInstancesResponse struct {
	// 修改实例的结果。
	Results *[]BatchRestartOrDeleteInstanceRespResults `json:"results,omitempty"`
}

func (o BatchRestartOrDeleteInstancesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchRestartOrDeleteInstancesResponse", string(data)}, " ")
}
