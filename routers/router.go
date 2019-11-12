package routers

import (
	"commons/http/crontroller/admin"
	"commons/http/crontroller/wx"
	"fmt"
	yaa "github.com/betacraft/yaag/gin"
	"github.com/gin-gonic/gin"
)

func InitRouter(port int, path string, node int64) *Engine {
	r := &Engine{gin.New(), adminc.NewAdminHttpAdminHandler(path, node), wx.NewWxHttpHandler(path, node)}
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	//yaag.Init(&yaag.Config{On: true, DocTitle: "Gin", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "", "Staging": ""}})
	r.Use(yaa.Document())

	r.weiXin()
	r.Engine.Run(fmt.Sprintf(":%d", port))
	return r
}
