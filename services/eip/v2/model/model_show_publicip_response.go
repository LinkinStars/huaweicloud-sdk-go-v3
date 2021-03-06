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
type ShowPublicipResponse struct {
	Publicip *PublicipShowResp `json:"publicip,omitempty"`
}

func (o ShowPublicipResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPublicipResponse", string(data)}, " ")
}
