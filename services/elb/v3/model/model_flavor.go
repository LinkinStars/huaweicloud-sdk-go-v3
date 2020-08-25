/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 负载均衡器规格信息
type Flavor struct {
	// 规格ID。
	Id   string      `json:"id"`
	Info *FlavorInfo `json:"info"`
	// 规格名称。
	Name string `json:"name"`
	// 共享。
	Shared bool `json:"shared"`
	// 项目ID。包括flavor所属的项目ID及其共享型的项目ID。
	ProjectId string `json:"project_id"`
	// L4和L7 分别表示四层和七层flavor。查询支持按type过滤
	Type string `json:"type"`
	// availability_zone_list字段，标志ELB对应L7-flavor在对应可用区是否可以售卖。 默认为[\"ALL\"]，所有可用区可售卖，若仅部分可用区可售卖则返回[\"cn-north-1a\",\"cn-north-1b\"]，若全部可不卖则为[]
	AvailabilityZoneIds []string `json:"availability_zone_ids,omitempty"`
}

func (o Flavor) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Flavor", string(data)}, " ")
}