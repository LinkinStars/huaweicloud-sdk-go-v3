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
type ShowCoordinatorsResponse struct {
	// 所有消费组对应的协调器列表。
	Coordinators *[]ShowCoordinatorsRespCoordinators `json:"coordinators,omitempty"`
}

func (o ShowCoordinatorsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowCoordinatorsResponse", string(data)}, " ")
}
