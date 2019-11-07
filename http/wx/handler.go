package wx

import (
	"commons/logic/wx"
)

type HttpWxHandler struct {
	logic wx.LogicHandler
	Node  int64
}

func NewWxHttpHandler(path string, node int64) *HttpWxHandler {
	return &HttpWxHandler{logic: wx.NewLogic(path), Node: node}
}
