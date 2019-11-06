package wx

import "commons/pkg/config"

type Config struct {
	Db  *db          `json:"db"`
	Rdx config.Redis `json:"rdx"`
	Es  config.Es    `json:"es"`
}

type db struct {
	Admin config.Config
}
