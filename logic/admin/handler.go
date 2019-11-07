package admin

import (
	"commons/dao/admin"
	"context"
)

type LogicHandler interface {
	handler
}

type handler interface {
	Exist(ctx context.Context, bean interface{}) bool
}
type Logic struct {
	db admin.DbHandler
}

func (l Logic) Exist(ctx context.Context, bean interface{}) bool {
	return l.db.Exist(ctx, bean)
}

var _ LogicHandler = Logic{}

func NewLogic(path string) LogicHandler {
	return &Logic{db: admin.NewDb(path)}
}
