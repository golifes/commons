package entiyWx

type Page struct {
	Ps int `json:"ps"  binding:"required"`
	Pn int `json:"pn"  binding:"required"`
}

type BizContent struct {
	ArticleId    string `json:"article_id" binding:"required"`
	Content      string `json:"content" binding:"required"`
	ContentStyle string `json:"content_style" binding:"required"`
	NickName     string `json:"nick_name"`
}

type RetList struct {
	OwId       int64  `json:"ow_id"`       //公号id
	Title      string `json:"title"`       //标题
	ContentUrl string `json:"content_url"` //url
	Ptime      int64  `json:"ptime"`       //发布时间
}

type Queue struct {
	Url       string `json:"url"`
	ArticleId string `json:"article_id"`
}

type Id struct {
	Id string `json:"id" binding:"required"`
}
