package wx

import (
	entiyWx "commons/entiy/wx"
	"commons/model/wx"
	"commons/pkg/app"
	"commons/pkg/e"
	"commons/tools/snowflake"
	"commons/tools/utils"
	"fmt"
	"net/http"
	"strings"
)

/**
页面添加微信号
params:{"wx_id":"","name":"","url":"","desc":"","biz":""}
*/
func (h HttpWxHandler) AddWx(ctx app.GContext) {
	var p entiyWx.WeiXin
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	exist := h.logic.Exist(g.NewContext(ctx), &wx.WeiXin{Biz: p.Biz})
	if exist {
		g.Json(http.StatusOK, e.WxExist, p.Name)
		return
	}
	p.Id = snowflake.NewNode(h.Node)
	p.Forbid = 1
	err = h.logic.Insert(g.NewContext(ctx), p)

	if err != nil {
		g.Json(http.StatusOK, e.Errors, p.Name)
	} else {
		g.Json(http.StatusOK, e.Success, p.Name)
	}
}

func (h HttpWxHandler) UpdateWxKey(ctx app.GContext) {

	/**
	如果biz是空,则是万能key
	*/

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

func (h HttpWxHandler) FindWxBiz(ctx app.GContext) {
	var p entiyWx.Wx
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	ps, pn := utils.Pagination(p.Ps, p.Pn, 10)
	query := []string{" forbid = ? "}
	values := []interface{}{1}
	if p.Name != "" {
		query = append(query, " and `name` like ? ")
		values = append(values, fmt.Sprintf("%s%s%s", "%", p.Name, "%"))
	}

	if p.Biz != "" {
		query = append(query, " and biz = ? ")
		values = append(values, p.Biz)
	}

	if p.Id != 0 {
		query = append(query, " and id = ? ")
		values = append(values, p.Id)
	}

	weiXin := make([]entiyWx.WxBiz, 0)

	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, "wei_xin", "id desc ", query, values, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)
}

func (h HttpWxHandler) ForBidWx(ctx app.GContext) {
	var p entiyWx.ForBidWx
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}
	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), wx.WeiXin{Forbid: 1}, []string{"forbid", "mtime"}, []string{"id = ?"}, []interface{}{p.Id})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Id)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}

func (h HttpWxHandler) WxList(ctx app.GContext) {
	var p entiyWx.WxList
	g, err := h.common(ctx, &p)
	if err != nil {
		return
	}

	ps, pn := utils.Pagination(p.Ps, p.Pn, 10)
	query := []string{" forbid = ? "}
	values := []interface{}{1}

	if p.Title != "" {
		query = append(query, " and `title` like ? ")
		values = append(values, fmt.Sprintf("%s%s%s", "%", p.Title, "%"))
	}

	if p.OwId != 0 {
		query = append(query, " and OwId = ? ")
		values = append(values, p.OwId)
	}

	if p.StartTime.Unix() > 0 {
		query = append(query, " and ptime >= ? ")
		values = append(values, p.StartTime)
	}

	if p.EndTime.Unix() > 0 {
		query = append(query, " and ptime <= ? ")
		values = append(values, p.StartTime)
	}
	oderBy := "ptime desc "
	if p.OrderBy != "" {
		oderBy = p.OrderBy
	}
	weiXin := make([]entiyWx.WeiXinList, 0)
	//
	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, "list", oderBy, query, values, ps, pn)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)
}

//获取多条数据
func (h HttpWxHandler) FindBizUinKey(ctx app.GContext) {
	//不接受任何参数
	var p entiyWx.Ps
	g, err := h.common(ctx, &p)

	if err != nil {
		return
	}

	if p.Ps == 0 {
		p.Ps = 10
	}
	weiXin := make([]entiyWx.WeiXinKey, 0)

	//这里需要改下，获取最近2小时未被抓取的公号
	list, count := h.logic.FindOne(g.NewContext(ctx), &weiXin, "wei_xin", "ctime desc ", nil, nil, p.Ps, 1)
	m := make(map[string]interface{})
	m["count"] = count
	m["list"] = list
	g.Json(http.StatusOK, e.Success, m)

}

func (h HttpWxHandler) AddWxList(ctx app.GContext) {
	g := app.G{ctx}

	var p entiyWx.List
	code := e.Success
	if !utils.CheckError(ctx.ShouldBindJSON(&p), "addList") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}

	var w wx.List
	//todo 先查询ArticleId是否存在，存在就不入库，不存在就存在入库
	w.ArticleId = p.ArticleId
	if h.logic.Exist(g.NewContext(ctx), &w) {
		g.Json(http.StatusOK, e.Errors, "")
		return
	}
	err := h.logic.Insert(g.NewContext(ctx), p)
	if !utils.CheckError(err, "insert addList") {
		code = e.Errors
	}
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpWxHandler) AddDetail(ctx app.GContext) {
	g := app.G{ctx}

	var p entiyWx.List
	code := e.Success
	if !utils.CheckError(ctx.ShouldBindJSON(&p), "addList") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	p.ContentStyle = strings.ReplaceAll(p.ContentStyle, "data-src", "src")
	b := h.logic.InsertEs(p.ArticleId, p)
	if !b {
		code = e.Errors
	}
	g.Json(http.StatusOK, code, "")
	return
}

func (h HttpWxHandler) SpiderRun(ctx app.GContext) {
	g := app.G{ctx}

	var p entiyWx.SpiderTime
	code := e.Success
	if !utils.CheckError(ctx.ShouldBindJSON(&p), "UpdateSpiderTime") {
		code = e.ParamError
		g.Json(http.StatusOK, code, "")
		return
	}
	cols := []string{}
	if p.Num == 0 {
		cols = append(cols, "stime", "run", "msg")
	} else {
		cols = append(cols, "stime", "num", "run", "msg")
	}
	affect, err := h.logic.UpdateStruct(g.NewContext(ctx), wx.WeiXin{Stime: p.Stime, Run: false, Num: p.Num, Msg: p.Msg}, cols, []string{" biz = ? "}, []interface{}{p.Biz})
	if !utils.CheckError(err, affect) {
		g.Json(http.StatusOK, e.UpdateWxError, p.Biz)
	} else {
		g.Json(http.StatusOK, e.Success, affect)
	}
}
