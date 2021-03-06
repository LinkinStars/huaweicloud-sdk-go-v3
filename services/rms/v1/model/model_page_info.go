/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 分页对象
type PageInfo struct {
	// 当前页数量
	CurrentCount *int32 `json:"current_count,omitempty"`
	// 下一页地址marker
	NextMarker *string `json:"next_marker,omitempty"`
}

func (o PageInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
