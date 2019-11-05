package admin

import "commons/pkg/config"

type Config struct {
	DB *db
}

type db struct {
	Admin config.Config
}
