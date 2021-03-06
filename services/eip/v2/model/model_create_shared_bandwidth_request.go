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

// Request Object
type CreateSharedBandwidthRequest struct {
	Body *CreateSharedBandwidhRequestBody `json:"body,omitempty"`
}

func (o CreateSharedBandwidthRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSharedBandwidthRequest", string(data)}, " ")
}
