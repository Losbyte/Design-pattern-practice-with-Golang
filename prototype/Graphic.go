package prototype

//此处为原型模式的实践代码

// Cloneable 原型接口
type Cloneable interface {
	Clone() Cloneable
}

// Graphic 图形接口
type Graphic interface {
	Cloneable
	Draw()
	GetID() string
}

