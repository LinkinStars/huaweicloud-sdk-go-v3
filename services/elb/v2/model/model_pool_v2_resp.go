/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// 后端云服务器组响应体
type PoolV2Resp struct {
	// 后端云服务器组的ID
	Id string `json:"id"`
	// 后端云服务器组所在的项目ID。
	ProjectId string `json:"project_id"`
	// 后端云服务器组所在的项目ID。
	TenantId string `json:"tenant_id"`
	// 后端云服务器组的名称。
	Name string `json:"name"`
	// 后端云服务器组的描述信息
	Description string `json:"description"`
	// 后端云服务器组的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp bool `json:"admin_state_up"`
	// 后端云服务器组绑定的负载均衡器ID的列表。
	Loadbalancers []ResourceList `json:"loadbalancers"`
	// 后端云服务器组关联的监听器ID的列表。
	Listeners []ResourceList `json:"listeners"`
	// 后端云服务器组关联的后端云服务器ID的列表。
	Members []ResourceList `json:"members"`
	// 后端云服务器组关联的健康检查的ID。
	HealthmonitorId    string              `json:"healthmonitor_id"`
	SessionPersistence *SessionPersistence `json:"session_persistence"`
	// 后端云服务器组的后端协议。
	Protocol PoolV2RespProtocol `json:"protocol"`
	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法。当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。
	LbAlgorithm PoolV2RespLbAlgorithm `json:"lb_algorithm"`
}

func (o PoolV2Resp) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"PoolV2Resp", string(data)}, " ")
}

type PoolV2RespProtocol struct {
	value string
}

type PoolV2RespProtocolEnum struct {
	UDP  PoolV2RespProtocol
	TCP  PoolV2RespProtocol
	HTTP PoolV2RespProtocol
}

func GetPoolV2RespProtocolEnum() PoolV2RespProtocolEnum {
	return PoolV2RespProtocolEnum{
		UDP: PoolV2RespProtocol{
			value: "UDP",
		},
		TCP: PoolV2RespProtocol{
			value: "TCP",
		},
		HTTP: PoolV2RespProtocol{
			value: "HTTP",
		},
	}
}

func (c PoolV2RespProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *PoolV2RespProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PoolV2RespLbAlgorithm struct {
	value string
}

type PoolV2RespLbAlgorithmEnum struct {
	ROUND_ROBIN       PoolV2RespLbAlgorithm
	LEAST_CONNECTIONS PoolV2RespLbAlgorithm
	SOURCE_IP         PoolV2RespLbAlgorithm
}

func GetPoolV2RespLbAlgorithmEnum() PoolV2RespLbAlgorithmEnum {
	return PoolV2RespLbAlgorithmEnum{
		ROUND_ROBIN: PoolV2RespLbAlgorithm{
			value: "ROUND_ROBIN",
		},
		LEAST_CONNECTIONS: PoolV2RespLbAlgorithm{
			value: "LEAST_CONNECTIONS",
		},
		SOURCE_IP: PoolV2RespLbAlgorithm{
			value: "SOURCE_IP",
		},
	}
}

func (c PoolV2RespLbAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *PoolV2RespLbAlgorithm) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}