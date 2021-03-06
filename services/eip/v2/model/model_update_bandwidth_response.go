/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateBandwidthResponse struct {
	Bandwidth *BandwidthResp `json:"bandwidth,omitempty"`
}

func (o UpdateBandwidthResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateBandwidthResponse", string(data)}, " ")
}
