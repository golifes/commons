package admin

import (
	"commons/pkg/config"
	"context"
	"github.com/go-redis/redis/v7"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
)

type Dao struct {
	c      Config
	engine *xorm.EngineGroup
	rdx    *redis.Client
	es     *elastic.Client
	//esIndex string
}

func (d Dao) Exist(ctx context.Context, model interface{}) bool {
	panic("implement me")
}

func (d Dao) Delete(ctx context.Context, id int64, model interface{}) (int64, error) {
	panic("implement me")
}

func (d Dao) FindOne(ctx context.Context, model interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	panic("implement me")
}

func (d Dao) GetOne(ctx context.Context, model interface{}, cols ...string) interface{} {
	panic("implement me")
}

func (d Dao) UpdateStruct(ctx context.Context, model interface{}, cols, query []string, values []interface{}) (int64, error) {
	panic("implement me")
}

func (d Dao) UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	panic("implement me")
}

func (d Dao) FindMany(ctx context.Context, bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64) {
	panic("implement me")
}

func (d Dao) Insert(ctx context.Context, beans ...interface{}) error {
	panic("implement me")
}

func (d Dao) DeleteMany(beans [][2]interface{}) error {
	panic("implement me")
}

func NewDb(path string) *Dao {
	//load 结构体绑定
	var c Config
	config.Load(path, &c)
	//es, index := config.LoadElastic(c.Es)
	return &Dao{engine: config.NewDb(c.Db.Admin), rdx: config.LoadRedis(c.Rdx)}
}
