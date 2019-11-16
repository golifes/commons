package wx

import (
	entiyWx "commons/entiy/wx"
	"commons/model/wx"
	"commons/pkg/app"
	"commons/pkg/e"
	"commons/tools/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

/**
提交到redis的任务
*/
const (
	StringQueue = "WxStringQueue"
)

func (h HttpWxHandler) AddQueue(ctx app.GContext) {

	/**
	接受一个列表推送
	*/

	var p []entiyWx.List
	code := e.Success

	g, err := h.common(ctx, &p)
	if err != nil {
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
	g.Json(http.StatusOK, code, "")
}

func (h HttpWxHandler) ErrorQueue(ctx app.GContext) {
	var p entiyWx.Queue
	code := e.Success
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	marshal, _ := json.Marshal(p)
	queue := h.logic.SetQueue(StringQueue, string(marshal))
	g.Json(http.StatusOK, code, queue)
}

/**
从set集合队列中获取数据
params: {"ps":必填，int "pn":必填,0}
*/
func (h HttpWxHandler) GetQueue(ctx app.GContext) {
	var p entiyWx.Page
	code := e.Success

	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	ps, _ := utils.Pagination(p.Ps, p.Pn, 10)
	pop := h.logic.SPop(StringQueue, int64(ps))
	m := make([]entiyWx.Queue, 0)
	for _, v := range pop {
		var q entiyWx.Queue
		_ = json.Unmarshal([]byte(v), &q)
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

	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	toMap := utils.StructToMap(p)

	h.logic.UpdateEs(p.ArticleId, toMap)
	g.Json(http.StatusOK, code, "")

}

func (h HttpWxHandler) GetList(ctx app.GContext) {

	var p entiyWx.Page
	code := e.Success

	g, err := h.common(ctx, &p)
	if err != nil {
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

	var p entiyWx.Id
	code := e.Success

	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	es := h.logic.GetOneEs(p.Id, []string{"content"}...)
	g.Json(http.StatusOK, code, es)
}

/**
获取抓取公号数据 每次请求返回2条数据
*/
func (h HttpWxHandler) GetWxBizList(ctx app.GContext) {

	var p entiyWx.BizList
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	ps, pn := utils.Pagination(p.Ps, 1, 10)

	weiXin := make([]entiyWx.WxBiz, 0)
	//unix := utils.StrTimeToUnix()
	query := []string{}
	if p.Incr {
		query = append(query, "incr = 1 and forbid = ? and  stime <= ? and run != 1 ")
	} else {
		query = append(query, " (forbid = ? and  stime <= ? and run != 1 )", " or msg !='ok' ")
	}
	values := []interface{}{1, p.Stime}

	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, "wei_xin", "id desc ", query, values, ps, pn)

	_values := []interface{}{}
	if count != 0 {
		xin := list.(*[]entiyWx.WxBiz)
		for _, v := range *xin {
			_values = append(_values, v.Biz)
		}
		//这里调用更新方法
		_count := len(_values)
		s := utils.JoinS(_count)
		affect, err := h.logic.UpdateStruct(g.NewContext(ctx), wx.WeiXin{Run: true}, []string{"run"}, []string{" biz in ( " + s + ")"}, _values)
		fmt.Println(affect, err)

	}

	m := make(map[string]interface{}, 0)
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)

}

func (h HttpWxHandler) UpdateBizKey(ctx app.GContext) {
	var p entiyWx.WeiXinKey
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	cols := []string{"key", "uin"}
	var query []string
	if p.Biz != "biz" {
		query = append(query, " biz = ? ")
	}

	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), wx.WeiXin{Key: p.Key, Uin: p.Uin}, cols, query, []interface{}{p.Biz})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Biz)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}
