package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T){
	order := NewOrder()

	fmt.Println("=== 订单状态流程测试 ===")
	fmt.Println("当前状态:", order.CurrentState())

	fmt.Println("尝试取消新建订单")
	order.Cancel()
	fmt.Println("当前状态:", order.CurrentState())

	fmt.Println("=== 新订单 ===")
	order = NewOrder()
	fmt.Println("当前状态:", order.CurrentState())

	fmt.Println("确认订单")
	order.Confirm()
	fmt.Println("当前状态:", order.CurrentState())

	fmt.Println("尝试发货")
	order.Ship()
	fmt.Println("当前状态:", order.CurrentState())

	fmt.Println("尝试收货")
	order.Deliver()
	fmt.Println("当前状态:", order.CurrentState())

	fmt.Println("=== 异常操作测试 ===")
	fmt.Println("尝试取消已收货订单")
	order.Cancel()

	fmt.Println("=== 状态模式解决的问题 ===")
	fmt.Println("1. 消除庞大的条件语句")
	fmt.Println("2. 使状态转换更加清晰")
	fmt.Println("3. 局部化状态相关行为")
	fmt.Println("4. 易于添加新状态")
}
