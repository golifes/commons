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

type Queue struct {
	Url  string `json:"url"`
	Biz  string `json:"biz"`
	OwId int    `json:"ow_id"`
}
