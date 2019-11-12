package routers

import "commons/http/middleware"

func (e *Engine) weiXin() {
	r := e.Group("/wx")
	{

		r.POST("/add", e.AddWx)       //添加公众号
		r.POST("/key", e.UpdateWxKey) //更新公众号key
		r.GET("/biz", e.FindWxBiz)    //获取公号列表数据
		r.POST("/forbid", e.ForBidWx) //禁用微信
		r.GET("/list", e.WxList)
		r.GET("/key", e.FindBizUinKey)
		//r.GET("/getBiz", e.FindBizUinKey)
		r.POST("/list", e.AddWxList)
		r.POST("/detail", e.AddDetail)
		r.POST("/spiderTime", e.UpdateSpiderTime)
		//差一个入库详情页和列表数据接口
	}
	v2 := e.Group("/api/v2/wx")
	{

		//提交列表数据
		v2.POST("/queue", e.AddQueue, middleware.API()) //提交队列任务
		v2.GET("/queue", e.GetQueue)                    //获取队列任务
		v2.POST("/detail", e.UpdateBizContent)          //更新详情数据
		v2.GET("/list", e.GetList)                      //获取前端列表数据
		v2.GET("/detail", e.GetOne)                     //获取详情数据

	}
}
