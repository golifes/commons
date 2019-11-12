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
	if !utils.CheckErrorArgs("add queue", ctx.ShouldBindJSON(&p)) {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	//这里需要优化，支持批量操作
	/**
	SetQueue支持批量操作
	*/
	for _, v := range p {
		articleId := v.ArticleId
		//存在
		if memList := h.logic.SAdd(articleId); memList == nil {
			continue
		}
		var queue entiyWx.Queue
		queue.Url = v.ContentURL
		queue.ArticleId = v.ArticleId
		marshal, _ := json.Marshal(queue)
		h.logic.SetQueue(StringQueue, string(marshal))
		h.logic.InsertEs(articleId, v)
	}
}

/**
从set集合队列中获取数据
params: {"ps":必填，int "pn":必填,0}
*/
func (h HttpWxHandler) GetQueue(ctx app.GContext) {
	g := app.G{ctx}
	var p entiyWx.Page
	code := e.Success
	if !utils.CheckErrorArgs("get queue", ctx.ShouldBindJSON(&p)) {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	ps, _ := utils.Pagination(p.Ps, p.Pn, 10)
	pop := h.logic.SPop(StringQueue, int64(ps))
	m := make([]entiyWx.Queue, 0)
	for _, v := range pop {
		var q entiyWx.Queue
		json.Unmarshal([]byte(v), &q)
		m = append(m, q)
	}
	g.Json(http.StatusOK, code, m)
	return
}

/**
更新详情数据
*/

func (h HttpWxHandler) UpdateBizContent(ctx app.GContext) {
	g := app.G{ctx}

	var p entiyWx.BizContent
	code := e.Success
	if !utils.CheckErrorArgs("get queue", ctx.ShouldBindJSON(&p)) {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	toMap := utils.StructToMap(p)

	h.logic.UpdateEs(p.ArticleId, toMap)
	g.Json(http.StatusOK, code, "")

}

func (h HttpWxHandler) GetList(ctx app.GContext) {
	g := app.G{ctx}

	var p entiyWx.Page
	code := e.Success
	if !utils.CheckErrorArgs("get queue", ctx.ShouldBindJSON(&p)) {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	ps, pn := utils.Pagination(p.Ps, p.Pn, 10)

	listEs, count := h.logic.GetListEs(ps, pn, []string{"article_id", "ow_id", "title", "content_url", "ptime"}...)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = listEs
	g.Json(http.StatusOK, code, m)
}

func (h HttpWxHandler) GetOne(ctx app.GContext) {
	g := app.G{ctx}

	var p entiyWx.Id
	code := e.Success
	if !utils.CheckErrorArgs("get queue", ctx.ShouldBindJSON(&p)) {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	es := h.logic.GetOneEs(p.Id, []string{"content"}...)
	g.Json(http.StatusOK, code, es)
}
