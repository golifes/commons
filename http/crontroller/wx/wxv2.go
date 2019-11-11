package wx

import (
	entiyWx "commons/entiy/wx"
	"commons/pkg/app"
	"commons/pkg/e"
	"commons/tools/utils"
	"encoding/json"
	"net/http"
)

/**
提交到redis的任务
*/
const (
	StringQueue = "WxStringQueue"
)

func (h HttpWxHandler) AddQueue(ctx app.GContext) {

	g := app.G{ctx}

	/**
	接受一个列表推送
	*/

	var p []entiyWx.List
	code := e.Success
	if !utils.CheckError(ctx.ShouldBindJSON(&p), "addList") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	for _, v := range p {
		articleId := v.ArticleId
		//存在
		if memList := h.logic.SAdd(articleId); memList == nil {
			continue
		}
		marshal, _ := json.Marshal(v)
		h.logic.Set(StringQueue, marshal)
		h.logic.InsertEs(articleId, v)
	}
	/**
	入redis  key articleId
			value: url

	其余入es

	更新es 用articleId去更新
	*/

}
