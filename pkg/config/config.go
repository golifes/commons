package config

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic/v7"
	"github.com/xormplus/xorm"
	"log"
)

const sqlConn = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"

type Config struct {
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

type Redis struct {
	Dns      string
	Pwd      string
	PoolSize int
	MaxIdle  int
	Db       int
}

type Es struct {
	Host  string //es host
	Index string // es index
}

func NewDb(c Config) *xorm.EngineGroup {
	conn := make([]string, 0)
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

	engine, err := xorm.NewEngineGroup("mysql", conn)
	//
	if err != nil {
		fmt.Println("err ", engine.Ping())
		panic(err)
	}
	//
	engine.SetMaxIdleConns(c.MaxIdle)
	engine.SetMaxOpenConns(c.MaxOpen)
	engine.ShowSQL(c.Show)
	return engine

}

func LoadRedis(c Redis) (rdx *redis.Client) {
	rdx = redis.NewClient(&redis.Options{
		Addr:     c.Dns,
		Password: c.Pwd, // no password set
		DB:       c.Db,
	})
	if _, err := rdx.Ping().Result(); err != nil {
		fmt.Println("c--->", c, err)
	}
	return
}

func LoadElastic(s Es) (es *elastic.Client, index string) {
	var err error
	index = s.Index
	fmt.Println(s)
	es, err = elastic.NewSimpleClient(elastic.SetURL(s.Host))
	if err != nil {
		log.Printf("elastic conn is error %s", err)
	}
	return
}
