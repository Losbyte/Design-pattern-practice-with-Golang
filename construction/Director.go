package construction

// Director 指挥者
type Director struct {
	builder ComputerBuilder
}

func NewDirector(builder ComputerBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) ConstructGamingPC() *Computer {
	return d.builder.
		SetCPU("Intel i9-13900K").
		SetMemory("32GB DDR5").
		SetStorage("1TB NVMe SSD").
		SetGPU("NVIDIA RTX 4090").
		Build()
}

func (d *Director) ConstructOfficePC() *Computer {
	return d.builder.
		SetCPU("Intel i5-12400").
		SetMemory("16GB DDR4").
		SetStorage("512GB SSD").
		Build() // 不设置GPU，使用集成显卡
}
