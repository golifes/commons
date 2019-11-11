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
		//提交列表数据{ 一个大json{biz content content_style}}  直接update es

		//提交数据到队列(json(列表页数据) ,入redis，协诚入es)
		/**
		redis :  article_id : url(队列),biz OwId

		*/
		v2.POST("/queue", e.AddQueue, middleware.API())
		//提交异常列表url(队列回填)(redis)

		//获取列表数据(es)

		//获取详情数据(es)
	}
}
