package main

import (
	"backend/initialize"
	"backend/server"
)

func main() {
	// init
	initialize.Init()
	// run router
	server.StartServer()
}
