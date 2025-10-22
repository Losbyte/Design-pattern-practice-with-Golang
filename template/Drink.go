package template

import "fmt"

// Beverage 饮料抽象类（模板类）
type Beverage interface {
	PrepareRecipe() // 模板方法
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
	CustomerWantsCondiments() bool // 钩子方法
}

// PrepareBeverage 准备饮料的模板方法
func PrepareBeverage(b Beverage) {
	fmt.Println("===== 开始制作饮料 =====")
	b.BoilWater()
	b.Brew()
	b.PourInCup()
	if b.CustomerWantsCondiments() {
		b.AddCondiments()
	}
	fmt.Println("===== 饮料制作完成 =====")
}
