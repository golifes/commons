package admin

import (
	"context"
)

type DbHandler interface {
	Handler
}

type Handler interface {
	Exist(ctx context.Context, model interface{}) bool
}

func (d Dao) Exist(ctx context.Context, model interface{}) bool {
	return false
}

var _ DbHandler = Dao{}
