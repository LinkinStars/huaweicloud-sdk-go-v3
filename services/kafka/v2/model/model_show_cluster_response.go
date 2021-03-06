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
type ShowClusterResponse struct {
	Cluster *ShowClusterRespCluster `json:"cluster,omitempty"`
}

func (o ShowClusterResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowClusterResponse", string(data)}, " ")
}
