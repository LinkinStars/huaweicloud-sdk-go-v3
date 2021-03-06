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
type ListProjectBugStaticsV4Response struct {
	// bug统计
	BugStatistics *[]BugStatisticResponseV4 `json:"bug_statistics,omitempty"`
}

func (o ListProjectBugStaticsV4Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectBugStaticsV4Response", string(data)}, " ")
}
