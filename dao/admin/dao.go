package admin

import (
	"commons/pkg/config"
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

func NewDb(path string) *Dao {
	//load 结构体绑定
	var c Config
	config.Load(path, &c)
	es, index := config.LoadElastic(c.Es)
	return &Dao{engine: config.NewDb(c.Db.Admin), rdx: config.LoadRedis(c.Rdx), es: es, esIndex: index}
}

//func NewDb(path string) *Dao {
//	return &Dao{config.NewConfig(path), config.NewDb()}
//}
