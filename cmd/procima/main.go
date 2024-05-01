package main

import (
	"flag"
	"fmt"
	appPkg "github.com/SShlykov/procima/internal/bootstrap/app"
	"os"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config", "path to configuration file")

	app, err := appPkg.New()
	if err != nil {
		fmt.Printf("failed to create app: %s\n", err.Error())
		os.Exit(2)
	}

	if err = app.Run(); err != nil {
		fmt.Printf("failed to run app: %s\n", err.Error())
		os.Exit(2)
	}
}
