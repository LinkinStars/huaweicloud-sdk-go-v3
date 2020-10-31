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
type ListPoolsResponse struct {
	// 后端云服务器对象组列表
	Pools *[]PoolV2Resp `json:"pools,omitempty"`
}

func (o ListPoolsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPoolsResponse", string(data)}, " ")
}