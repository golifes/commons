package admin

import (
	"commons/pkg/config"
	"github.com/xormplus/xorm"
)

type Dao struct {
	c      Config
	engine *xorm.EngineGroup
}

func NewPath(path string) (d *Dao) {
	//load 结构体绑定
	var c Config
	config.Load(path, &c)
	d = &Dao{engine: config.NewDb(c.DB.Admin)}
	return
}
