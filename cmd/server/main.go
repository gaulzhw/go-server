package main

import (
	"os"

	"github.com/gaulzhw/go-server/cmd/server/app"
)

func main() {
	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
