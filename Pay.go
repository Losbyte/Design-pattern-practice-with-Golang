package strategy

import "fmt"

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(amount float64) error
	GetName() string
}

// PaymentContext 支付上下文
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) ExecutePayment(amount float64) error {
	fmt.Printf("使用 %s 支付: ¥%.2f\n", p.strategy.GetName(), amount)
	return p.strategy.Pay(amount)
}
