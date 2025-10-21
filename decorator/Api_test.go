package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T){
	// 订一杯浓缩咖啡，不加配料
	espresso := &Espresso{}
	fmt.Printf("%s $%.2f\n", espresso.GetDescription(), espresso.Cost())

	// 订一杯深焙咖啡，加双份摩卡和奶泡
	var darkRoast Beverage = &DarkRoast{}
	darkRoast = NewMocha(darkRoast)
	darkRoast = NewMocha(darkRoast)
	darkRoast = NewWhip(darkRoast)
	fmt.Printf("%s $%.2f\n", darkRoast.GetDescription(), darkRoast.Cost())

	// 订一杯浓缩咖啡，加牛奶、摩卡和奶泡
	var beverage Beverage = &Espresso{}
	beverage = NewMilk(beverage)
	beverage = NewMocha(beverage)
	beverage = NewWhip(beverage)
	fmt.Printf("%s $%.2f\n", beverage.GetDescription(), beverage.Cost())
}
