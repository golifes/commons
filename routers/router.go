package routers

import (
	"commons/http/admin"
	"commons/http/wx"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string, path string) *Engine {
	r := &Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path), wx.NewWxHttpHandler(path), func() string {
		return port
	},
	}
	r.Engine.Run(":8080")
	return r
}
