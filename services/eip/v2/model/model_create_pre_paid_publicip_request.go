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
type CreatePrePaidPublicipRequest struct {
	Body *CreatePrePaidPublicipRequestBody `json:"body,omitempty"`
}

func (o CreatePrePaidPublicipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePrePaidPublicipRequest", string(data)}, " ")
}
