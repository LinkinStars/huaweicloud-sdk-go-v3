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
type ListPublicipTagsRequest struct {
}

func (o ListPublicipTagsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPublicipTagsRequest", string(data)}, " ")
}
