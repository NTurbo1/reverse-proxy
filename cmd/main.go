package main

import (
	"fmt"

	"github.com/nturbo1/reverse-proxy/configs"
	"github.com/nturbo1/reverse-proxy/internal/log"
)

func main() {
	appConfigs, err := configs.InitAppConfigs()
	if err != nil {
		fmt.Printf("Failed to initialize app configs due to: %s\n", err)
		return
	}

	fmt.Println("Initializing the logger...")
	log.InitLogger(appConfigs.LogLevel)
	log.Info("Configurations: %s", appConfigs)
}
