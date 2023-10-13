package nonecode

import (
	"errors"
	"net/http"
)

// Guard represents an authentication entry
type MiddlewareAuth struct {
	JWT *JWT
}

// NewGuard returns a new guard
func NewMiddlewareAuth(jwt *JWT) *MiddlewareAuth {
	return &MiddlewareAuth{
		JWT: jwt,
	}
}

// AuthJWT checks the jwt token is authenticated
func (g *MiddlewareAuth) AuthJWT(c *Context) {
	jwtToken := c.GetHeader(g.JWT.AuthorizationKey)
	jwt, ok := g.JWT.Verify(jwtToken)
	if !ok {
		c.FailWithMsg(http.StatusUnauthorized, "AccessToken error")
		c.Abort()
		return
	}
	scope, _ := c.Get(AUTH_SCOPE)
	if scope.(string) == jwt.Scope {
		setJwtWithContext(c, jwt)
		c.Next()
	} else {
		c.FailWithMsg(http.StatusUnauthorized, "AccessToken permission error")
		c.Abort()

	}
	// if (c.GetString("api_permission") == "user" && jwt.Scope == model.ScopeUser) ||
	// 	(c.GetString("api_permission") == "admin" && jwt.Scope == model.ScopeUserAdmin) ||
	// 	(c.GetString("api_permission") == "customer" && (jwt.Scope == model.ScopeUserAdmin || jwt.Scope == model.ScopeUserCustomer)) {

	// 	// set value
	// 	setJwtWithContext(c, jwt)
	// 	c.Next()
	// } else {
	// 	c.FailWithMsg(http.StatusUnauthorized, "AccessToken permission error")
	// 	c.Abort()
	// }
}

// setJwtWithContext 设置gin context信息
func setJwtWithContext(ctx *Context, jwtToken *Claims) {
	ctx.Set("authinfo", jwtToken)
	// ctx.Set(model.ContextValueID, jwtToken.ID)
	// ctx.Set(model.ContextValueScope, jwtToken.Scope)
	// ctx.Set(model.ContextValueUserName, jwtToken.Username)
}

// GetContextValueScope 从gin context 中获取权限范围
func GetContextValueScope(ctx *Context) string {
	// return ctx.GetString(model.ContextValueScope)
	return ctx.GetString("")
}

// // GetContextValueUserName 从gin context中获取用户名称
// func GetContextValueUserName(ctx *gin.Context) string {
// 	return ctx.GetString(model.ContextValueUserName)
// }

func (ctx *Context) GetAuth() *Claims {
	// if ctx.GetString(model.ContextValueScope) == model.ScopeUser {
	// 	return ctx.GetUint(model.ContextValueID)
	// }
	authinfo, ok := ctx.Get("authinfo")
	if !ok {
		return nil
	}
	return authinfo.(*Claims)
}

// func GetMgmtAdminAuthID(ctx *nonecode.Context) uint {
// 	if ctx.GetString(model.ContextValueScope) == model.ScopeUserAdmin {
// 		return ctx.GetUint(model.ContextValueID)
// 	}
// 	return 0
// }

// func GetMgmtCustomerAuthID(ctx *nonecode.Context) uint {
// 	if ctx.GetString(model.ContextValueScope) == model.ScopeUserCustomer {
// 		return ctx.GetUint(model.ContextValueID)
// 	}
// 	return 0
// }

func (c *Context) SetToken(user *JwtUser) error {
	if c == nil {
		return errors.New("nonecode.Context empty")
	}
	jwt, _ := c.Get("JWT")
	tk, err := jwt.(*JWT).NewAccessToken(user)
	if err != nil {
		return err
	}
	c.Header("", tk)
	return nil
}
