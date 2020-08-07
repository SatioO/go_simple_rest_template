package main

import (
	"strconv"

	"github.com/satioO/go_practices/app"
)

func main() {
	server := app.Server{}
	server.Initialize()
	port := strconv.Itoa(server.Config.Server.Port)
	server.Run(":" + port)
}
