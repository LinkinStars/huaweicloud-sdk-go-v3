/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 镜像标签
type ImageTag struct {
	// 标签key值
	Key *string `json:"key,omitempty"`
	// 标签value值
	Value *string `json:"value,omitempty"`
}

func (o ImageTag) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ImageTag", string(data)}, " ")
}
