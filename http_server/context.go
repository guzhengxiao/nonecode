package http_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type ApiResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg" example:"错误原因"`
}

//	func (c *Context) DB() *gorm.DB {
//		return DB.Session(&gorm.Session{
//			PrepareStmt: true,
//			// NowFunc: func() time.Time {
//			// 	return time.Now().Local()time.LoadLocation(lib.ConfigData.App.Timezone)
//			// },
//		})
//	}
func (c *Context) Succ(d interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: 200,
		Data: d,
		Msg:  "Success",
	})
}

func (c *Context) FailWithMsg(code int, msg string) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: code,
		Data: nil,
		Msg:  msg,
	})
}

func (c *Context) ResJson(code int, d interface{}, msg string) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: code,
		Data: d,
		Msg:  msg,
	})
}
