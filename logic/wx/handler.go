package wx

import (
	"commons/dao/wx"
	"context"
)

type LogicHandler interface {
	handler
}

type handler interface {
	Exist(ctx context.Context, bean ...interface{}) bool
	//删除
	Delete(ctx context.Context, id int64, bean interface{}) (int64, error)
	//单表查询
	FindOne(ctx context.Context, bean interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64)
	//获取一条记录
	GetOne(ctx context.Context, bean interface{}, cols ...string) interface{}
	//以结构体的方式更新
	UpdateStruct(ctx context.Context, bean interface{}, cols, query []string, values []interface{}) (int64, error)
	//以map的方式更新
	UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error)
	//多表查询
	Join2Table(ctx context.Context, bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64)
	Insert(ctx context.Context, beans ...interface{}) error //事物
	Delete2Table(beans [][2]interface{}) error              //事物
}
type Logic struct {
	db wx.DbHandler
}

func (l Logic) Exist(ctx context.Context, bean ...interface{}) bool {
	return l.db.Exist(ctx, bean)
}

func (l Logic) Delete(ctx context.Context, id int64, bean interface{}) (int64, error) {
	return l.db.Delete(ctx, id, bean)
}

func (l Logic) FindOne(ctx context.Context, bean interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	return l.db.FindOne(ctx, bean, table, orderBy, query, values, ps, pn)
}

func (l Logic) GetOne(ctx context.Context, bean interface{}, cols ...string) interface{} {
	return l.db.GetOne(ctx, bean, cols...)
}

func (l Logic) UpdateStruct(ctx context.Context, bean interface{}, cols, query []string, values []interface{}) (int64, error) {
	return l.db.UpdateStruct(ctx, bean, cols, query, values)
}

func (l Logic) UpdateMap(ctx context.Context, table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	return l.db.UpdateMap(ctx, table, m, cols, query, values)
}

func (l Logic) Join2Table(ctx context.Context, bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64) {
	return l.db.Join2Table(ctx, bean, table, alias, cols, orderBy, ps, pn, query, values, join)
}

func (l Logic) Insert(ctx context.Context, beans ...interface{}) error {
	panic("implement me")
}

func (l Logic) Delete2Table(beans [][2]interface{}) error {
	panic("implement me")
}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{db: wx.NewDb(path)}
}
