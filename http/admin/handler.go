package adminc

import "commons/logic/admin"

type HttpAdminHandler struct {
	logic admin.LogicHandler
	Node  int64
}

func NewAdminHttpAdminHandler(path string, node int64) *HttpAdminHandler {
	return &HttpAdminHandler{logic: admin.NewLogic(path), Node: node}
}
