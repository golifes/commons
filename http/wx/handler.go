package wx

import (
	"commons/logic/wx"
)

type HttpWxHandler struct {
	logic wx.LogicHandler
}

func NewWxHttpHandler(path string) *HttpWxHandler {
	return &HttpWxHandler{logic: wx.NewLogic(path)}
}
