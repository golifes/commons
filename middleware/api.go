package middleware

import (
	"commons/pkg/app"
	"commons/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
给别人提供api的白名单
*/

func API() gin.HandlerFunc {
	return func(c app.GContext) {

		g := app.G{c}
		ip := g.ClientIP()
		ipList := [...]string{"39.108.190.186"}
		f := false
		for _, v := range ipList {
			if ip == v {
				f = true
				break
			}
		}
		if !f {
			g.Json(http.StatusOK, e.Forbid, "")
			g.Abort()
			return
		}
		g.Next()
	}
}
