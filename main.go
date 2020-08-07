package main

import (
	"github.com/satioO/go_practices/app"
)

func main() {
	server := app.Server{}
	server.Initialize()
	server.Run(":3000")
}
