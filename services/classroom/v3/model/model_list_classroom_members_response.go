/*
 * Classroom
 *
 * devcloud classedge api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListClassroomMembersResponse struct {
	// 课堂成员列表
	Members *[]ClassroomMember `json:"members,omitempty"`
	// 课堂成员总数
	Total *int32 `json:"total,omitempty"`
}

func (o ListClassroomMembersResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListClassroomMembersResponse", string(data)}, " ")
}
