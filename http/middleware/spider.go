package middleware

import (
	"commons/pkg/app"
	"commons/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Spider() gin.HandlerFunc {
	return func(c app.GContext) {

		//ipList := [...]string{"219.142.7.99", "127.0.0.1", "58.87.64.219", "180.168.177.4"}
		if !m.BlackList(BlackListSpider, ip) {
			g.Json(http.StatusOK, e.Forbid, "")
			g.Abort()
			return
		}

		g.Next()
	}
}
