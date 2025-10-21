package chain

import "fmt"

// AuthHandler 认证处理器
type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Handle(req *Request) {
	if token, ok := req.Headers["Authorization"]; ok && token == "valid-token" {
		fmt.Println("AuthHandler: Authentication successful")
		if a.next != nil {
			a.next.Handle(req)
		}
	} else {
		fmt.Println("AuthHandler: Authentication failed - invalid token")
	}
}

// LoggingHandler 日志处理器
type LoggingHandler struct {
	BaseHandler
}

func (l *LoggingHandler) Handle(req *Request) {
	fmt.Printf("LoggingHandler: %s %s\n", req.Method, req.Path)
	if l.next != nil {
		l.next.Handle(req)
	}
}

// RoutingHandler 路由处理器
type RoutingHandler struct {
	BaseHandler
}

func (r *RoutingHandler) Handle(req *Request) {
	switch req.Path {
	case "/api/users":
		fmt.Println("RoutingHandler: Handling user API request")
	case "/api/products":
		fmt.Println("RoutingHandler: Handling product API request")
	default:
		fmt.Println("RoutingHandler: 404 Not Found")
		return // 不继续传递
	}

	if r.next != nil {
		r.next.Handle(req)
	}
}

// CacheHandler 缓存处理器
type CacheHandler struct {
	BaseHandler
	cache map[string]string
}

func NewCacheHandler() *CacheHandler {
	return &CacheHandler{
		cache: make(map[string]string),
	}
}

func (c *CacheHandler) Handle(req *Request) {
	if req.Method == "GET" {
		if body, exists := c.cache[req.Path]; exists {
			fmt.Printf("CacheHandler: Serving from cache for %s\n", req.Path)
			req.Body = body
			return // 从缓存返回，不继续传递
		}
	}

	if c.next != nil {
		c.next.Handle(req)

		// 如果是GET请求且没有缓存，则缓存结果
		if req.Method == "GET" && req.Body != "" {
			c.cache[req.Path] = req.Body
			fmt.Printf("CacheHandler: Caching response for %s\n", req.Path)
		}
	}
}


