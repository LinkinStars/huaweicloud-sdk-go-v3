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
type NeutronDeleteFloatingIpResponse struct {
}

func (o NeutronDeleteFloatingIpResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"NeutronDeleteFloatingIpResponse", string(data)}, " ")
}
