package entiyWx

import "time"

type WeiXin struct {
	Id         int64     `json:"id"`                       //主键Id
	WxId       string    `json:"wx_id" binding:"required"` //微信id
	Name       string    `json:"name" binding:"required"`  //微信名称
	Url        string    `json:"url" binding:"required"`   //微信头像
	Desc       string    `json:"desc" binding:"required"`  //公号描述
	Biz        string    `json:"biz" binding:"required"`   //公号biz
	Count      string    `json:"count"`                    //公号文章数量
	Forbid     int       `json:"forbid" xorm:"default 1"`  //是否被禁用
	Key        string    `json:"key"`                      //公号Key
	Uin        string    `json:"uin"`                      //用户唯一标识
	Ctime      time.Time `json:"ctime" xorm:"created"`     //创建时间
	Mtime      time.Time `json:"mtime" xorm:"updated"`     //更新时间
	SpiderTime time.Time `json:"spider_time"`              //最后一次抓取时间
	Note       string    `json:"note"`                     //备用字段
}

type WeiXinParams struct {
	Id       int    `json:"id"`       //查询biz
	Keywords string `json:"keywords"` //关键字
	From     string `json:"from"`     //发布时间起始
	To       string `json:"to"`       // 发布时间截止
	Title    string `json:"title"`    //标题模糊查询
	Pn       int    `json:"pn"`
	Ps       int    `json:"ps"`
	Type     int    `json:"type"` //导出当前页还是导出全部数据  1是导出全部  -1导出当前页
}

type WeiXinKey struct {
	Id   int64  `json:"id"`
	WxId string `json:"wx_id"`
	Biz  string `json:"biz" binding:"required"`
	Key  string `json:"key" binding:"required"`
	Uin  string `json:"uin" binding:"required"`
}

type Wx struct {
	Id   int64  `json:"id"`   //通过id查询
	Name string `json:"name"` //通过name模糊查询
	Biz  string `json:"biz"`  //通过biz查询
	Ps   int
	Pn   int
}

type WxBiz struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Biz  string `json:"biz"`
	WxId string `json:"wx_id"`
	Uin  string `json:"uin"`
	Key  string `json:"key"`
}

type ForBidWx struct {
	Id string `json:"id" binding:"required"`
}

//获取微信文章列表数据
type WxList struct {
	Id        int64     `json:"id"`                      //公号id
	ArticleId string    `json:"article_id"`              //文章id
	Title     string    `json:"title"`                   //标题
	Forbid    int       `json:"forbid" xorm:"default 1"` //是否被禁用
	StartTime time.Time `json:"start_time"`              //发布时间
	EndTime   time.Time `json:"end_time"`                //结束时间
	Ps        int       `json:"ps"`                      //分页
	Pn        int       `json:"pn"`                      //分页
	OrderBy   string    `json:"order_by"`                //排序
}

type Ps struct {
	Ps int `json:"ps"`
}

type ParamsAddWxList struct {
	Id    int64  `json:"id"`
	Url   string `json:"url"  binding:"required" `   //文章url
	Title string `json:"title"  binding:"required" ` //文章标题
	Ptime int64  `json:"ptime"`                      //发布时间
}
