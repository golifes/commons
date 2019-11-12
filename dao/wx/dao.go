package wx

import (
	"commons/pkg/config"
	"context"
	"github.com/go-redis/redis/v7"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
)

type Dao struct {
	c       Config
	engine  *xorm.EngineGroup
	rdx     *redis.Client
	es      *elastic.Client
	esIndex string
}

func (d Dao) GetListEs(ps, pn int, cols ...string) ([]interface{}, interface{}) {
	return d.getListEs(ps, pn, cols...)
}

func (d Dao) UpdateEs(id string, m map[string]interface{}) bool {
	return d.updateEs(id, m)
}

func (d Dao) GetOneEs(id string, cols ...string) interface{} {
	return d.getOneEs(id, cols...)
}

func (d Dao) SPop(key string, ps int64) []string {
	return d.sPop(key, ps)
}

func (d Dao) SetQueue(key string, members ...interface{}) bool {
	return d.setQueue(key, members...)
}

func (d Dao) SAdd(members ...interface{}) (memList []string) {
	return d.sAdd(members...)
}

func (d Dao) Set(key string, value interface{}) bool {
	return d.set(key, value, 7*24*60)
}

func (d Dao) Get(key string) []byte {
	return d.get(key)
}

func (d Dao) InsertEs(id string, bean interface{}) bool {
	return d.addEsOne(id, bean)
}

func NewDb(path string) *Dao {
	//load 结构体绑定
	var c Config
	config.Load(path, &c)
	es, index := config.LoadElastic(c.Es)
	return &Dao{engine: config.NewDb(c.Db.Admin), rdx: config.LoadRedis(c.Rdx), es: es, esIndex: index}
}

func (d Dao) Exist(ctx context.Context, bean ...interface{}) bool {
	return d.exist(bean)
}

func (d Dao) Delete(ctx context.Context, id int64, bean interface{}) (int64, error) {
	return d.delete(id, bean)
}

func (d Dao) FindOne(ctx context.Context, bean interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	return d.findOne(bean, table, orderBy, query, values, ps, pn)
}

func (d Dao) GetOne(ctx context.Context, bean interface{}, cols ...string) interface{} {
	return d.getOne(bean)
}

func (d Dao) UpdateStruct(ctx context.Context, bean interface{}, cols, query []string, values []interface{}) (int64, error) {
	return d.updateStruct(bean, cols, query, values)
}

func (d Dao) UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	return d.updateMap(table, m, cols, query, values)
}

func (d Dao) Join2Table(ctx context.Context, bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64) {
	return d.join2Table(bean, table, alias, cols, orderBy, ps, pn, query, values, join)
}

func (d Dao) Insert(ctx context.Context, beans ...interface{}) error {
	return d.insertMany(beans...)
}

func (d Dao) Delete2Table(beans [][2]interface{}) error {
	return d.delete2Table(beans)
}
