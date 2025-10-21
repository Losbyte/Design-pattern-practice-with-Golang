package state


import "fmt"

// OrderState 订单状态接口
type OrderState interface {
	Confirm(o *Order)
	Cancel(o *Order)
	Ship(o *Order)
	Deliver(o *Order)
}

// Order 订单上下文
type Order struct {
	state OrderState
}

func NewOrder() *Order {
	return &Order{state: &NewState{}}
}

func (o *Order) SetState(state OrderState) {
	o.state = state
}

func (o *Order) Confirm() {
	o.state.Confirm(o)
}

func (o *Order) Cancel() {
	o.state.Cancel(o)
}

func (o *Order) Ship() {
	o.state.Ship(o)
}

func (o *Order) Deliver() {
	o.state.Deliver(o)
}

func (o *Order) CurrentState() string {
	return fmt.Sprintf("%T", o.state)
}



