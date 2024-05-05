package main

import (
	"flag"
	"fmt"
	appPkg "github.com/SShlykov/procima/procima/internal/bootstrap/app"
	"os"
)

const BadCode = 2

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config", "path to configuration file")

	app, err := appPkg.New(configPath)
	if err != nil {
		fmt.Printf("failed to create app: %s\n", err.Error())
		os.Exit(BadCode)
	}

	if err = app.Run(); err != nil {
		fmt.Printf("failed to run app: %s\n", err.Error())
		os.Exit(BadCode)
	}
}
