package proxy

import "fmt"

// Image 图片接口
type Image interface {
	Display()
}

// RealImage 真实图片类
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	image := &RealImage{filename: filename}
	image.loadFromDisk() // 模拟高开销的加载操作
	return image
}

func (i *RealImage) Display() {
	fmt.Printf("Displaying image: %s\n", i.filename)
}

func (i *RealImage) loadFromDisk() {
	fmt.Printf("Loading image from disk: %s (expensive operation)\n", i.filename)
}
