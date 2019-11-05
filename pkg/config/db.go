package config

import (
	"fmt"
	"github.com/xormplus/xorm"
)

/**
这里是数据库连接函数
"Host": "47.75.105.53:9200",
    "Port": 3306,
    "User": "root",
    "Db": "fadmin",
    "Pwd": "abc123456",
    "Show": true,
    "MaxOpen": 256,
    "MaxIdle": 150
*/
const sqlConn = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"

type Config struct {
	Addr  string
	Write struct {
		Addr string
		Port int
		User string
		Pwd  string
		Db   string
	}
	Read []struct {
		Addr string
		Port int
		User string
		Pwd  string
		Db   string
	}
	Show    bool
	MaxOpen int
	MaxIdle int
}

func NewDb(c Config) *xorm.EngineGroup {
	fmt.Println(c)
	conn := make([]string, 1)
	w := fmt.Sprintf(sqlConn,
		c.Write.User,
		c.Write.Pwd,
		c.Write.Addr,
		c.Write.Db)

	conn = append(conn, w)
	for _, v := range c.Read {
		s := fmt.Sprintf(sqlConn,
			v.User,
			v.Pwd,
			v.Addr,
			v.Db)
		conn = append(conn, s)
	}
	//
	engine, _ := xorm.NewEngineGroup("mysql", c.Read)
	//
	//if err != nil || engine.Ping() != nil {
	//	panic(err)
	//}
	//
	//engine.SetMaxIdleConns(c.Db.MaxIdle)
	//engine.SetMaxOpenConns(c.Db.MaxOpen)
	//engine.ShowSQL(c.Db.Show)
	return engine

}
