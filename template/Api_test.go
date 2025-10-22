package template

import (
	"fmt"
	"testing"
)

func TestTea(t *testing.T) {
	fmt.Println("=== 模板方法模式演示：饮料制作 ===")

	// 制作加调料的茶
	teaWithCondiments := NewTea(true)
	teaWithCondiments.PrepareRecipe()

	// 制作不加调料的茶
	teaWithoutCondiments := NewTea(false)
	teaWithoutCondiments.PrepareRecipe()

	// 制作加调料的咖啡
	coffeeWithCondiments := NewCoffee(true)
	coffeeWithCondiments.PrepareRecipe()

	// 制作不加调料的咖啡
	coffeeWithoutCondiments := NewCoffee(false)
	coffeeWithoutCondiments.PrepareRecipe()

	fmt.Println("=== 模板方法模式解决的问题 ===")
	fmt.Println("1. 代码复用：茶和咖啡共享相同的准备流程")
	fmt.Println("2. 流程控制：父类控制算法流程")
	fmt.Println("3. 灵活性：通过钩子方法控制是否添加调料")
	fmt.Println("4. 减少重复代码：避免子类重复实现相同步骤")
}
