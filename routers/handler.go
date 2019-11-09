package routers

import (
	"commons/http/admin"
	"commons/http/wx"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*adminc.HttpAdminHandler
	*wx.HttpWxHandler
}
