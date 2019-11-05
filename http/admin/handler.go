package adminc

import "commons/logic/admin"

type HttpAdminHandler struct {
	logic admin.LogicHandler
}

func NewAdminHttpAdminHandler(path string) *HttpAdminHandler {
	return &HttpAdminHandler{logic: admin.NewLogic(path)}
}
