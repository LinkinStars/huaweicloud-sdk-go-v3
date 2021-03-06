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

// 用户昵称
type UpdateUserNickNameRequestV4 struct {
	// 用户昵称
	NickName string `json:"nick_name"`
}

func (o UpdateUserNickNameRequestV4) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdateUserNickNameRequestV4", string(data)}, " ")
}
