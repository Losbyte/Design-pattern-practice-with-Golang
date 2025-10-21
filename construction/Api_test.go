package construction

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T){
	builder := NewStandardComputerBuilder()
	myPC := builder.
		SetCPU("AMD Ryzen 7 5800X").
		SetMemory("16GB DDR4").
		SetStorage("1TB SSD").
		SetGPU("AMD RX 6700 XT").
		Build()
	fmt.Println(myPC.Show())

	// 方式2：使用指挥者
	director := NewDirector(NewStandardComputerBuilder())
	gamingPC := director.ConstructGamingPC()
	fmt.Println(gamingPC.Show())

	officePC := director.ConstructOfficePC()
	fmt.Println(officePC.Show())
}
