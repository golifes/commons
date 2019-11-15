package wx

/**
白名单
*/

func (h HttpWxHandler) BlackList(key string, ip string) bool {
	return h.logic.SisMember(key, ip)
}
