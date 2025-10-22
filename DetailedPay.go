package strategy

import "fmt"

type CreditCardPayment struct {
	cardNumber string
	cvv        string
}

func NewCreditCardPayment(cardNumber, cvv string) *CreditCardPayment {
	return &CreditCardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (c *CreditCardPayment) Pay(amount float64) error {
	// 模拟信用卡支付逻辑
	fmt.Printf("信用卡支付: 卡号%s, 金额¥%.2f\n", c.cardNumber[:4]+"****"+c.cardNumber[12:], amount)
	fmt.Println("验证CVV码...")
	fmt.Println("支付成功!")
	return nil
}

func (c *CreditCardPayment) GetName() string {
	return "信用卡"
}

// AlipayPayment 支付宝支付策略
type AlipayPayment struct {
	account string
}

func NewAlipayPayment(account string) *AlipayPayment {
	return &AlipayPayment{account: account}
}

func (a *AlipayPayment) Pay(amount float64) error {
	// 模拟支付宝支付逻辑
	fmt.Printf("支付宝支付: 账户%s, 金额¥%.2f\n", a.account, amount)
	fmt.Println("跳转支付宝页面...")
	fmt.Println("用户确认支付...")
	fmt.Println("支付成功!")
	return nil
}

func (a *AlipayPayment) GetName() string {
	return "支付宝"
}

// WechatPayment 微信支付策略
type WechatPayment struct {
	openID string
}

func NewWechatPayment(openID string) *WechatPayment {
	return &WechatPayment{openID: openID}
}

func (w *WechatPayment) Pay(amount float64) error {
	// 模拟微信支付逻辑
	fmt.Printf("微信支付: OpenID%s, 金额¥%.2f\n", w.openID[:8]+"****", amount)
	fmt.Println("调起微信支付...")
	fmt.Println("用户输入密码...")
	fmt.Println("支付成功!")
	return nil
}

func (w *WechatPayment) GetName() string {
	return "微信支付"
}
