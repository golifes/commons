package routers

import (
	"commons/http/admin"
	"commons/http/wx"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(port int, path string, node int64) *Engine {
	r := &Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path, node), wx.NewWxHttpHandler(path, node)}
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.weiXin()
	r.Engine.Run(fmt.Sprintf(":%d", port))
	return r
}
