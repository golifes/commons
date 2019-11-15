package main

import (
	"commons/routers/router"
	"flag"
)

func main() {
	//admin.NewPath("config/config.json")
	path := flag.String("-c", "config/config.json", "config.conf")
	port := flag.Int("-p", 8080, "port")
	node := flag.Int64("-n", 1, "node")
	router.InitRouter(*port, *path, *node)
}
