package router

import (
	"commons/http/admin"
	"commons/http/wx"
	"commons/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(port int, path string, node int64) *routers.Engine {
	r := &routers.Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path, node), wx.NewWxHttpHandler(path, node)}
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	weiXin(r)
	r.Engine.Run(fmt.Sprintf(":%d", port))

	return r
}
