// Code generated by goctl. DO NOT EDIT.
package types

type TaskOrderRequest struct {
	Computation float64 `form:"computation"`
}

type TaskOrderResponse struct {
	Ack string `json:"ack"`
}