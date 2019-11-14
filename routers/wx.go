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
		//r.POST("/spiderTime", e.UpdateSpiderTime)
		//差一个入库详情页和列表数据接口
	}

	//自己用
	v2 := e.Group("/api/v2/wx")
	v2.Use(middleware.Spider())
	{
		//获取biz和key等信息 这里接手一个时间,获取每个时间段的
		v2.GET("/biz", e.GetWxBizList)
		//更新微信key  如果biz是biz，则是万能key
		v2.POST("/biz", e.UpdateBizKey)

		v2.POST("/run", e.SpiderRun)

		//提交列表数据
		v2.POST("/queue", e.AddQueue)          //提交队列任务
		v2.GET("/queue", e.GetQueue)           //获取队列任务
		v2.POST("/detail", e.UpdateBizContent) //更新详情数据
		v2.GET("/list", e.GetList)             //获取前端列表数据
		v2.GET("/detail", e.GetOne)            //获取详情数据
	}

	//给bean的api
	v3 := e.Group("/api/bean/wx")
	v3.Use(middleware.API())
	{
		v3.GET("/list", e.GetList, middleware.API())  //获取前端列表数据
		v3.GET("/detail", e.GetOne, middleware.API()) //获取详情数据
	}

}
