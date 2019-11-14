package wx

type List struct {
	ArticleId              string `json:"article_id" xorm:"article_id"` //主键id
	Biz                    string `json:"biz"`                          //
	Mid                    int64  `json:"mid"`                          //
	Idx                    int64  `json:"idx"`                          //
	OwId                   int64  `json:"ow_id"`                        //自定义微信id
	Oid                    int64  `json:"oid" xorm:"oid"`               //原始id,就是公号数据中的id
	AudioFileid            int64  `json:"audio_fileid" xorm:"audio_fileid"`
	Title                  string `json:"title"  xorm:"title"`
	ContentURL             string `json:"content_url" xorm:"content_url"`
	SourceURL              string `json:"source_url" xorm:"source_url"`
	Cover                  string `json:"cover" xorm:"cover"`
	Author                 string `json:"author" xorm:"author"`
	DelFlag                int64  `json:"del_flag" xorm:"del_flag"`
	PlayURL                string `json:"play_url" xorm:"play_url"`
	Content                string `json:"content" xorm:"content"`
	ItemShowType           int64  `json:"item_show_type" xorm:"item_show_type"`
	MaliciousTitleReasonID int64  `json:"malicious_title_reason_id" xorm:"malicious_title_reason_id"`
	Digest                 string `json:"digest" xorm:"digest"`
	Fileid                 int64  `json:"fileid" xorm:"fileid"`
	CopyrightStat          int64  `json:"copyright_stat" xorm:"copyright_stat"`
	Duration               int64  `json:"duration" xorm:"duration"`
	MaliciousContentType   int64  `json:"malicious_content_type" xorm:"malicious_content_type"`
	Subtype                int64  `json:"subtype" xorm:"subtype"`
	Type                   int64  `json:"type" xorm:"type"`
	Ptime                  int64  `json:"ptime" xorm:"ptime"`
	Fakeid                 string `json:"fakeid" xorm:"fakeid"`
	Status                 int64  `json:"status" xorm:"status"`
	Ctime                  int64  `json:"ctime" xorm:"created"` //创建时间
}

//func (w WeiXinList) Deadline() (deadline time.Time, ok bool) {
//	panic("implement me")
//}
//
//func (w WeiXinList) Done() <-chan struct{} {
//	panic("implement me")
//}
//
//func (w WeiXinList) Err() error {
//	panic("implement me")
//}
//
//func (w WeiXinList) Value(key interface{}) interface{} {
//	panic("implement me")
//}
//
