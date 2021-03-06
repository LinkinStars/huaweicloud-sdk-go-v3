/*
 * EPS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 配额信息
type QuotasDetail struct {
	// 资源配额
	Resources []EpQuotas `json:"resources"`
}

func (o QuotasDetail) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"QuotasDetail", string(data)}, " ")
}
