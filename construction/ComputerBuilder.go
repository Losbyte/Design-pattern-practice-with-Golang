package construction

// ComputerBuilder 建造者接口
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetStorage(storage string) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	Build() *Computer
}

// StandardComputerBuilder 具体建造者
type StandardComputerBuilder struct {
	computer *Computer
}

func NewStandardComputerBuilder() *StandardComputerBuilder {
	return &StandardComputerBuilder{computer: &Computer{}}
}

func (b *StandardComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *StandardComputerBuilder) SetMemory(memory string) ComputerBuilder {
	b.computer.Memory = memory
	return b
}

func (b *StandardComputerBuilder) SetStorage(storage string) ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *StandardComputerBuilder) SetGPU(gpu string) ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *StandardComputerBuilder) Build() *Computer {
	return b.computer
}
