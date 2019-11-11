package routers

import (
	"commons/http/crontroller/admin"
	"commons/http/crontroller/wx"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*adminc.HttpAdminHandler
	*wx.HttpWxHandler
}
