package middleware

import (
	"commons/http/wx"
	"commons/pkg/app"
	"commons/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	BlackListSpider = "BlackListSpider"
	BlackListApi    = "BlackListApi"
)

func Spider(m *wx.HttpWxHandler) gin.HandlerFunc {
	return func(c app.GContext) {
		g := app.G{c}
		ip := g.ClientIP()

		//ipList := [...]string{"219.142.7.99", "127.0.0.1", "58.87.64.219", "180.168.177.4"}
		if !m.BlackList(BlackListSpider, ip) {
			g.Json(http.StatusOK, e.Forbid, "")
			g.Abort()
			return
		}

		g.Next()
	}
}
