package http_server

import "github.com/gin-gonic/gin"

// type HandlerFunc func(c *Context)
type RouterGroup struct {
	*gin.RouterGroup
}

func (s *RouterGroup) Group(relativePath string, handlers ...func(c *Context)) *RouterGroup {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return &RouterGroup{s.RouterGroup.Group(relativePath, RHandles...)}
}

func (r *RouterGroup) Any(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.GET(relativePath, rHandles...)
}

// GET 拓展Get请求（子）
func (r *RouterGroup) GET(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.GET(relativePath, rHandles...)
}

// POST 拓展Post请求（子）
func (r *RouterGroup) POST(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.POST(relativePath, rHandles...)
}

// Use 拓展中间件注册
func (r *RouterGroup) Use(middlewares ...func(c *Context)) gin.IRoutes {
	rMiddlewares := make([]gin.HandlerFunc, 0)
	for _, middleware := range middlewares {
		rMiddlewares = append(rMiddlewares, handleFunc(middleware))
	}
	return r.RouterGroup.Use(rMiddlewares...)
}

func handleFunc(handler func(c *Context)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		handler(&Context{Context: c})
	}
}
