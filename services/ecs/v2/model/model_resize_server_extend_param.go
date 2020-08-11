/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type ResizeServerExtendParam struct {
	// 下单订购后，是否自动从客户的账户中支付，而不需要客户手动去进行支付。  - “true”：是（自动支付） - “false”：否（需要客户手动支付）  > 说明： >  > 当弹性云服务器是按包年包月计费时生效，该值为空时默认为客户手动支付。
	IsAutoPay string `json:"isAutoPay,omitempty"`
}