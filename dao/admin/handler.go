package admin

import "context"

type DbHandler interface {
	Handler
}

type Handler interface {
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

var _ DbHandler = Dao{}
