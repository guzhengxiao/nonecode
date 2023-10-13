package nonecode

import "gorm.io/gorm"

type HandlerFunc func(c *Context)
type MiddlewareGorm struct {
	*gorm.DB
}

func NewMiddlewareGorm(db *gorm.DB) *MiddlewareGorm {
	return &MiddlewareGorm{DB: db}
}

func (g *MiddlewareGorm) MiddlewareGorm(c *Context) {
	c.Set(MIDDLEWARE_GORM, g)
	c.Next()
}
