package chain

// Request HTTP请求结构
type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    string
}

// Handler 处理接口
type Handler interface {
	Handle(req *Request)
	SetNext(handler Handler)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}
