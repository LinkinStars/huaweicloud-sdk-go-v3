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
type ShowIssueV4Request struct {
	ProjectId string `json:"project_id"`
	IssueId   int32  `json:"issue_id"`
}

func (o ShowIssueV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowIssueV4Request", string(data)}, " ")
}
