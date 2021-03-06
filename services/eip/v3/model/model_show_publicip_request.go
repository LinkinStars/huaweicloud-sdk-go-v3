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
type ShowPublicipRequest struct {
	PublicipId string    `json:"publicip_id"`
	Fields     *[]string `json:"fields,omitempty"`
}

func (o ShowPublicipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPublicipRequest", string(data)}, " ")
}
