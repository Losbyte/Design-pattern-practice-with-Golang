package prototype

// GraphicManager 图形管理器（原型注册表）
type GraphicManager struct {
	prototypes map[string]Graphic
}

func NewGraphicManager() *GraphicManager {
	return &GraphicManager{
		prototypes: make(map[string]Graphic),
	}
}

// Register 注册原型
func (m *GraphicManager) Register(name string, proto Graphic) {
	m.prototypes[name] = proto
}

// CreateGraphic 根据名称创建图形副本
func (m *GraphicManager) CreateGraphic(name string) Graphic {
	proto, ok := m.prototypes[name]
	if !ok {
		return nil
	}
	return proto.Clone().(Graphic)
}
