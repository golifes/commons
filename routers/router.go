package routers

import (
	"commons/http/admin"
	"github.com/gin-gonic/gin"
)

func InitRouter(port string, path string) *Engine {
	r := &Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path), func() string {
		return port
	},
	}
	r.Engine.Run(":8080")
	return r
}
