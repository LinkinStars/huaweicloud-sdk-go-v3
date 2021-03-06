/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddMemberV4Request struct {
	ProjectId string              `json:"project_id"`
	Body      *AddMemberRequestV4 `json:"body,omitempty"`
}

func (o AddMemberV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"AddMemberV4Request", string(data)}, " ")
}
