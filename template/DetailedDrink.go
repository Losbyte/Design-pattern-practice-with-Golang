package template

import "fmt"

// Tea 茶类
type Tea struct {
	wantsCondiments bool // 控制钩子方法
}

func NewTea(wantsCondiments bool) *Tea {
	return &Tea{wantsCondiments: wantsCondiments}
}

func (t *Tea) PrepareRecipe() {
	PrepareBeverage(t)
}

func (t *Tea) BoilWater() {
	fmt.Println("煮沸水")
}

func (t *Tea) Brew() {
	fmt.Println("用沸水冲泡茶叶")
}

func (t *Tea) PourInCup() {
	fmt.Println("把茶倒入杯子")
}

func (t *Tea) AddCondiments() {
	fmt.Println("加柠檬")
}

func (t *Tea) CustomerWantsCondiments() bool {
	return t.wantsCondiments
}

// Coffee 咖啡类
type Coffee struct {
	wantsCondiments bool // 控制钩子方法
}

func NewCoffee(wantsCondiments bool) *Coffee {
	return &Coffee{wantsCondiments: wantsCondiments}
}

func (c *Coffee) PrepareRecipe() {
	PrepareBeverage(c)
}

func (c *Coffee) BoilWater() {
	fmt.Println("煮沸水")
}

func (c *Coffee) Brew() {
	fmt.Println("用沸水冲泡咖啡粉")
}

func (c *Coffee) PourInCup() {
	fmt.Println("把咖啡倒入杯子")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("加糖和牛奶")
}

func (c *Coffee) CustomerWantsCondiments() bool {
	return c.wantsCondiments
}
