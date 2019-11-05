package admin

import "commons/pkg/config"

/**
这里是当前项目db的配置
*/
type Config struct {
	Db  *db          `json:"db"`
	Rdx config.Redis `json:"rdx"`
	Es  config.Es    `json:"es"`
}

type db struct {
	Admin config.Config
}
