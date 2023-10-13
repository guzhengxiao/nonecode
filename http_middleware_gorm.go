package nonecode

import "gorm.io/gorm"

type HandlerFunc func(c *Context)
type GormMiddleware struct {
	*gorm.DB
}

func NewMiddlewareGorm(db *gorm.DB) *GormMiddleware {
	return &GormMiddleware{DB: db}
}

func (g *GormMiddleware) GormMiddleware() HandlerFunc {
	return func(c *Context) {
		c.Set(MIDDLEWARE_GORM, g)
		c.Next()
	}
}
