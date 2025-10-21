package proxy

import "fmt"

// ImageProxy 图片代理
type ImageProxy struct {
	filename    string
	realImage   *RealImage
	accessLevel int // 访问控制级别
}

func NewImageProxy(filename string, accessLevel int) *ImageProxy {
	return &ImageProxy{
		filename:    filename,
		accessLevel: accessLevel,
	}
}

func (p *ImageProxy) Display() {
	if p.accessLevel < 3 {
		fmt.Printf("Access denied: insufficient privileges for %s\n", p.filename)
		return
	}

	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename) // 延迟初始化
	}
	p.realImage.Display()
}

// SetAccessLevel 修改访问权限
func (p *ImageProxy) SetAccessLevel(level int) {
	p.accessLevel = level
}
