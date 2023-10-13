package nonecode

import (
	"github.com/guzhengxiao/nonecode/http_server"

	"github.com/gin-gonic/gin"
)

func New(e *gin.Engine) *http_server.HttpServer {
	return http_server.NewHttpServer(e)
}
