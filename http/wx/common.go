package wx

import (
	"commons/pkg/app"
	"commons/pkg/e"
	"commons/tools/utils"
	"errors"

	"net/http"
)

func (h HttpWxHandler) common(ctx app.GContext, obj interface{}) (app.G, error) {
	g := app.G{Context: ctx}

	code := e.Success
	bindJSON := ctx.ShouldBindJSON(obj)
	if !utils.CheckError(bindJSON, "bind params error") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return g, errors.New("参数绑定失败")
	}
	return g, nil
}
