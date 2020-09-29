package main

import (
	"fmt"
	"paas/global"
	"paas/initialize"
)

func main() {
	initialize.Mysql()
	initialize.DBTables()
	initialize.Docker()
	runServer()
}

func runServer() {
	Router := initialize.Router()
	Router.Static("source", "./source")

	address := fmt.Sprintf(":%d", global.GCONFIG.Addr)
	Router.Run(address)

}
