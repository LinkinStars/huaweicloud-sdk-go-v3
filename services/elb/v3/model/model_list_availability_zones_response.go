/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAvailabilityZonesResponse struct {
	// 可用区列表。  > 获取可用区集合列表后，在（如创建LB时）设置可用区，选择的多个可用区必须同时在同一个集合中。
	AvailabilityZones *[][]AvailabilityZone `json:"availability_zones,omitempty"`
	// 请求ID。  注：自动生成。
	RequestId *string `json:"request_id,omitempty"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
