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
type ChangeApplicationRequest struct {
	ApplicationId string             `json:"application_id"`
	Body          *ApplicationModify `json:"body,omitempty"`
}

func (o ChangeApplicationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeApplicationRequest", string(data)}, " ")
}
