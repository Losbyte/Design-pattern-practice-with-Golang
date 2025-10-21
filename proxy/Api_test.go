package proxy

import (
	"fmt"
	"testing"
)

func TestPhotoProxy(t *testing.T){
	fmt.Println("=== 代理模式演示 ===")

	// 创建代理对象（此时不会加载真实图片）
	image1 := NewImageProxy("photo1.jpg", 3)
	image2 := NewImageProxy("photo2.jpg", 2)
	image3 := NewImageProxy("photo3.jpg", 4)

	// 第一次显示（会触发加载）
	fmt.Println("第一次显示:")
	image1.Display() // 有权限，会加载并显示
	image2.Display() // 无权限，拒绝访问
	image3.Display() // 有权限，会加载并显示

	// 第二次显示（不会重新加载）
	fmt.Println("第二次显示:")
	image1.Display() // 已加载，直接显示
	image2.Display() // 仍然无权限

	// 修改权限后再次尝试
	fmt.Println("修改权限后:")
	image2.SetAccessLevel(3)
	image2.Display() // 现在有权限了，会加载并显示

	fmt.Println("=== 代理模式解决的问题 ===")
	fmt.Println("1. 延迟初始化：推迟高开销操作直到真正需要时")
	fmt.Println("2. 访问控制：根据权限决定是否允许访问")
	fmt.Println("3. 减少资源消耗：避免不必要的对象创建")
	fmt.Println("4. 增强安全性：增加额外的访问控制层")
}
