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
type ListDomainNotAddedProjectsV4Request struct {
	Offset *int32 `json:"offset,omitempty"`
	Limit  *int32 `json:"limit,omitempty"`
}

func (o ListDomainNotAddedProjectsV4Request) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListDomainNotAddedProjectsV4Request", string(data)}, " ")
}
