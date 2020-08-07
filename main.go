package main

import (
	root "github.com/satioO/go_practices/app"
)

func main() {
	server := root.App{}
	server.Initialize()
	server.SetupRoutes()
	server.Run(":3000")
}
