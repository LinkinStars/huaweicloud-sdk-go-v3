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

// Request Object
type CreateApplicationRequest struct {
	Body *ApplicationCreate `json:"body,omitempty"`
}

func (o CreateApplicationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateApplicationRequest", string(data)}, " ")
}
