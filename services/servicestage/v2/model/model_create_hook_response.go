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
type CreateHookResponse struct {
	// hook ID。
	Id *string `json:"id,omitempty"`
	// hook类型。
	Type *string `json:"type,omitempty"`
	// 回滚URL。
	CallbackUrl *string `json:"callback_url,omitempty"`
}

func (o CreateHookResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateHookResponse", string(data)}, " ")
}
