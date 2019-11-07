package routers

import (
	"commons/http/admin"
	"commons/http/wx"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string, path string, node int64) *Engine {
	r := &Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path, node), wx.NewWxHttpHandler(path, node), func() string {
		return port
	},
	}

	r.weiXin()
	r.Engine.Run(":8080")
	return r
}
