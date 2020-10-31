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

// 创建后端云服务器组的请求体
type CreatePoolV2Req struct {
	// 后端云服务器组的后端协议。取值：UDP、TCP、HTTP。当指定istener_id创建后端云服务器组时，后端云服务器组的protocol和它关联的监听器的protocol有如下关系：监听器的protocol为TCP时，后端云服务器组的protocol必须为TCP。监听器的protocol为UDP时，后端云服务器组的protocol必须为UDP。监听器的protocol为HTTP或TERMINATED_HTTPS时，后端云服务器组的protocol必须为HTTP。
	Protocol CreatePoolV2ReqProtocol `json:"protocol"`
	// 后端云服务器组的负载均衡算法，取值：ROUND_ROBIN：加权轮询算法；LEAST_CONNECTIONS：加权最少连接算法；SOURCE_IP：源IP算法；当该字段的取值为SOURCE_IP时，后端云服务器组绑定的后端云服务器的weight字段无效。
	LbAlgorithm string `json:"lb_algorithm"`
	// 后端云服务器组关联的负载均衡器ID。listener_id和loadbalancer_id中至少指定一个。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// 后端云服务器组关联的监听器的ID。listener_id和loadbalancer_id中至少指定一个。
	ListenerId *string `json:"listener_id,omitempty"`
	// 后端云服务器组所在的项目ID。
	TenantId *string `json:"tenant_id,omitempty"`
	// 后端云服务器组的名称。
	Name *string `json:"name,omitempty"`
	// 后端云服务器组的描述信息
	Description *string `json:"description,omitempty"`
	// 后端云服务器组的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp       *bool               `json:"admin_state_up,omitempty"`
	SessionPersistence *SessionPersistence `json:"session_persistence,omitempty"`
}

func (o CreatePoolV2Req) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePoolV2Req", string(data)}, " ")
}

type CreatePoolV2ReqProtocol struct {
	value string
}

type CreatePoolV2ReqProtocolEnum struct {
	UDP  CreatePoolV2ReqProtocol
	TCP  CreatePoolV2ReqProtocol
	HTTP CreatePoolV2ReqProtocol
}

func GetCreatePoolV2ReqProtocolEnum() CreatePoolV2ReqProtocolEnum {
	return CreatePoolV2ReqProtocolEnum{
		UDP: CreatePoolV2ReqProtocol{
			value: "UDP",
		},
		TCP: CreatePoolV2ReqProtocol{
			value: "TCP",
		},
		HTTP: CreatePoolV2ReqProtocol{
			value: "HTTP",
		},
	}
}

func (c CreatePoolV2ReqProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreatePoolV2ReqProtocol) UnmarshalJSON(b []byte) error {
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