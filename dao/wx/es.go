package wx

import (
	entiyWx "commons/entiy/wx"
	"commons/tools/utils"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

/**
es添加数据
*/

func (d Dao) addEsOne(id string, bean interface{}) bool {

	if bytes, err := json.Marshal(bean); utils.CheckError(err, bytes) {
		do, err := d.es.Index().Index(d.esIndex).Id(id).BodyString(string(bytes)).Do(context.Background())
		if utils.CheckError(err, do) {
			return true
		}
		return false
	} else {
		return false
	}
}

//批量入库
func (d Dao) AddBulk() {

}

func (d Dao) articles(detail entiyWx.WeiXinParams, ps, pn int) (interface{}, interface{}) {
	/**
	这里拼接es sql


	select a from table where forbid = 1 and ( title like "%aaa%" or text like "%bbb%")

	https://www.tuicool.com/articles/NFVzeqy
	*/
	//zuiwaiceng  query bool
	query := elastic.NewBoolQuery()
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewTermQuery("forbid", 1))
	query.Filter(boolQuery)
	filter := elastic.NewBoolQuery()
	newBoolQuery := elastic.NewBoolQuery()
	queryFilter := filter.Must(newBoolQuery)
	if detail.Title != "" {
		newBoolQuery.Should(elastic.NewWildcardQuery("title", detail.Title))
	}

	if detail.Keywords != "" {
		split := strings.Split(detail.Keywords, ",")
		for _, v := range split {
			newBoolQuery.Should(elastic.NewWildcardQuery("title", v))
			newBoolQuery.Should(elastic.NewWildcardQuery("text", v))
		}
	}
	//if detail.Id != 0 {
	//	biz := d.findBizById(detail.Id)
	//	if biz != "" {
	//		query.Must(elastic.NewMatchQuery("biz", biz))
	//	}
	//}
	const t = "2006-01-02 15:04:05"
	if detail.From != "" && detail.To != "" {
		query.Filter(elastic.NewRangeQuery("ptime").Gte(utils.Str2Time(t, detail.From)).Lte(utils.Str2Time(t, detail.To)))
	} else if detail.From != "" {
		query.Filter(elastic.NewRangeQuery("ptime").Gte(utils.Str2Time(t, detail.From)).Lte(time.Now()))
	}

	query.Filter(queryFilter)

	field := elastic.NewFetchSourceContext(true)
	field.Include("id", "text", "text_style", "biz", "author", "original", "word_cloud", "summary", "title", "forbid")
	source, _ := query.Source()
	utils.PrintQuery(source)
	result, err := d.es.Search().FetchSourceContext(field).Index(d.esIndex).Query(query).Do(context.Background())
	if utils.CheckError(err, result) {
		array := make([]interface{}, len(result.Hits.Hits))
		for index, hit := range result.Hits.Hits {
			//var r ret
			//json.Unmarshal(hit.Source, &r)
			array[index] = hit.Source
		}
		return array, result.Hits.TotalHits.Value
	}
	return nil, 0
}
