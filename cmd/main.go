package main

import (
	"commons/routers"
	"flag"
)

func main() {
	//admin.NewPath("config/config.json")
	path := flag.String("-c", "config/config.json", "config.conf")
	port := flag.String("-p", ":8080", "port")
	node := flag.Int64("-n", 1, "node")
	routers.InitRouter(*port, *path, *node)
}
