/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type KeyAliasInfo struct {
	// 密钥ID。
	KeyId *string `json:"key_id,omitempty"`
	// 密钥别名。
	KeyAlias *string `json:"key_alias,omitempty"`
}

func (o KeyAliasInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeyAliasInfo", string(data)}, " ")
}
