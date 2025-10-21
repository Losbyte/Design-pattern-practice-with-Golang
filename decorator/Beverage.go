package decorator

// Beverage 饮料接口
type Beverage interface {
	GetDescription() string
	Cost() float64
}

// Espresso 浓缩咖啡
type Espresso struct{}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) Cost() float64 {
	return 1.99
}

// DarkRoast 深焙咖啡
type DarkRoast struct{}

func (d *DarkRoast) GetDescription() string {
	return "Dark Roast Coffee"
}

func (d *DarkRoast) Cost() float64 {
	return 0.99
}