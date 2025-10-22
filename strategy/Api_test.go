package strategy

import (
	"fmt"
	"testing"
)

func TestPay(t *testing.T) {
	fmt.Println("=== 策略模式演示：支付系统 ===")

	// 创建支付策略
	creditCard := NewCreditCardPayment("1234567812345678", "123")
	alipay := NewAlipayPayment("user@alipay.com")
	wechat := NewWechatPayment("wx1234567890abcdef")

	// 创建支付上下文
	payment := NewPaymentContext(creditCard)

	// 使用不同策略进行支付
	amounts := []float64{100.50, 200.00, 350.75}

	for i, amount := range amounts {
		fmt.Printf("\n--- 第%d笔交易 ---\n", i+1)

		// 动态切换支付策略
		switch i {
		case 0:
			payment.SetStrategy(creditCard)
		case 1:
			payment.SetStrategy(alipay)
		case 2:
			payment.SetStrategy(wechat)
		}

		err := payment.ExecutePayment(amount)
		if err != nil {
			fmt.Printf("支付失败: %v\n", err)
		}
	}

}
