package state

import "fmt"

// NewState 新建状态
type NewState struct{}

func (n *NewState) Confirm(o *Order) {
	fmt.Println("订单确认成功")
	o.SetState(&PaidState{})
}

func (n *NewState) Cancel(o *Order) {
	fmt.Println("订单已取消")
	o.SetState(&CancelledState{})
}

func (n *NewState) Ship(o *Order) {
	fmt.Println("错误：新建订单不能直接发货")
}

func (n *NewState) Deliver(o *Order) {
	fmt.Println("错误：新建订单不能直接收货")
}

// PaidState 已支付状态
type PaidState struct{}

func (p *PaidState) Confirm(o *Order) {
	fmt.Println("订单已确认，无需重复操作")
}

func (p *PaidState) Cancel(o *Order) {
	fmt.Println("订单已取消，将退款")
	o.SetState(&CancelledState{})
}

func (p *PaidState) Ship(o *Order) {
	fmt.Println("订单已发货")
	o.SetState(&ShippedState{})
}

func (p *PaidState) Deliver(o *Order) {
	fmt.Println("错误：订单尚未发货，不能收货")
}

// ShippedState 已发货状态
type ShippedState struct{}

func (s *ShippedState) Confirm(o *Order) {
	fmt.Println("订单已确认，无需重复操作")
}

func (s *ShippedState) Cancel(o *Order) {
	fmt.Println("错误：已发货订单不能取消")
}

func (s *ShippedState) Ship(o *Order) {
	fmt.Println("订单已发货，无需重复操作")
}

func (s *ShippedState) Deliver(o *Order) {
	fmt.Println("订单已收货")
	o.SetState(&DeliveredState{})
}

// DeliveredState 已收货状态
type DeliveredState struct{}

func (d *DeliveredState) Confirm(o *Order) {
	fmt.Println("订单已完成，无需操作")
}

func (d *DeliveredState) Cancel(o *Order) {
	fmt.Println("错误：已收货订单不能取消")
}

func (d *DeliveredState) Ship(o *Order) {
	fmt.Println("错误：已收货订单不能发货")
}

func (d *DeliveredState) Deliver(o *Order) {
	fmt.Println("订单已完成，无需重复收货")
}

// CancelledState 已取消状态
type CancelledState struct{}

func (c *CancelledState) Confirm(o *Order) {
	fmt.Println("错误：已取消订单不能确认")
}

func (c *CancelledState) Cancel(o *Order) {
	fmt.Println("订单已取消，无需重复操作")
}

func (c *CancelledState) Ship(o *Order) {
	fmt.Println("错误：已取消订单不能发货")
}

func (c *CancelledState) Deliver(o *Order) {
	fmt.Println("错误：已取消订单不能收货")
}
