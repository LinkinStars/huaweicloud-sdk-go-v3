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

// Response Object
type BatchDeleteIssuesV4Response struct {
}

func (o BatchDeleteIssuesV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteIssuesV4Response", string(data)}, " ")
}
