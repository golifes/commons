package routers

import (
	"commons/http/admin"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	*gin.Engine
	*adminc.HttpAdminHandler
	Port func() string
}
