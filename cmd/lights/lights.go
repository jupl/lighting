package main

import (
	"github.com/joho/godotenv"
	"github.com/jupl/lighting/env"
	"github.com/jupl/lighting/light"
	"github.com/jupl/lighting/logger"
)

func main() {
	log := logger.New()
	errorLog := logger.NewError()

	// Attempt to read .env file
	if godotenv.Load() != nil {
		errorLog.Println("Cannot read .env file")
	}

	// Read light config
	config, err := env.LightConfig(false)
	if err != nil {
		errorLog.Fatalln(err)
	}
	if config.Host == "" {
		log.Println("HUE_LIGHT_HOST is not set, auto selecting host")
	}

	// Get information and display
	info, err := light.Info(config)
	if err != nil {
		errorLog.Fatalln(err)
	}
	log.Println(info)
}
