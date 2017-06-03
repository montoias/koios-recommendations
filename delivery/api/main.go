package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	mainConfig "github.com/montoias/koios-recommendations/config"
	deps "github.com/montoias/koios-recommendations/config/api"
)

const port = "8080"

func main() {
	// threads configuration
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	log.Printf("NumCPU: %d", numCPU)

	config, err := mainConfig.Init()
	if err != nil {
		panic(err)
	}

	runApplication(config)
}

func runApplication(config *mainConfig.MainConfig) {
	dependencies := deps.SetupDependencies(config)

	// global settings
	http.DefaultClient.Timeout = time.Duration(dependencies.DefaultTimeout) * time.Second

	app := dependencies.App
	app.Run(fmt.Sprintf("%s:%s", config.GetString("app.host"), port))
}
