package middleware

import (
	"commons/pkg/app"
	"commons/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Spider() gin.HandlerFunc {
	return func(c app.GContext) {

		g := app.G{c}
		ip := g.ClientIP()
		ipList := [...]string{"219.142.7.99", "127.0.0.1", "58.87.64.219"}
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
