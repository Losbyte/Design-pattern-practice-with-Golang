package decorator

// CondimentDecorator 配料装饰器
type CondimentDecorator struct {
	beverage Beverage
}

func (c *CondimentDecorator) GetDescription() string {
	return c.beverage.GetDescription()
}

func (c *CondimentDecorator) Cost() float64 {
	return c.beverage.Cost()
}

// Mocha 摩卡配料
type Mocha struct {
	CondimentDecorator
}

func NewMocha(b Beverage) *Mocha {
	return &Mocha{CondimentDecorator{beverage: b}}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.20
}

// Milk 牛奶配料
type Milk struct {
	CondimentDecorator
}

func NewMilk(b Beverage) *Milk {
	return &Milk{CondimentDecorator{beverage: b}}
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.beverage.Cost() + 0.10
}

// Whip 奶泡配料
type Whip struct {
	CondimentDecorator
}

func NewWhip(b Beverage) *Whip {
	return &Whip{CondimentDecorator{beverage: b}}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", Whip"
}

func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 0.15
}


