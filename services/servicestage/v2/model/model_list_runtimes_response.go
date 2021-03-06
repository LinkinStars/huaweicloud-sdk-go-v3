/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRuntimesResponse struct {
	// 运行时列表。
	Runtimes *[]RuntimeTypeView `json:"runtimes,omitempty"`
}

func (o ListRuntimesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRuntimesResponse", string(data)}, " ")
}
