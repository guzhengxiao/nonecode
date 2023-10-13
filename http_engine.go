package nonecode

import "github.com/gin-gonic/gin"

type HttpServer struct {
	*gin.Engine
	// DB *gorm.DB
}

func NewHttpServer(g *gin.Engine) *HttpServer {
	return &HttpServer{Engine: g}
}

// Group 重写路由组注册
func (s *HttpServer) Group(relativePath string, handlers ...func(c *Context)) *RouterGroup {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return &RouterGroup{s.Engine.Group(relativePath, RHandles...)}
}

// GET 拓展Get请求（根）
func (s *HttpServer) GET(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return s.Engine.GET(relativePath, RHandles...)
}

// POST 拓展POST请求（根）
func (s *HttpServer) POST(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return s.Engine.POST(relativePath, RHandles...)
}
