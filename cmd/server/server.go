package main

import (
	"os"

	"github.com/gaulzhw/go-server/internal/server"
)

// @title Apiserver Example API
// @version 1.0
// @description apiserver demo

// @host localhost:8080
// @BasePath /v1.
func main() {
	command := server.NewServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
