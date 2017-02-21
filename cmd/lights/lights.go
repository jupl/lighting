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

	// Validate light config
	lightConfig := env.LightConfig()
	if lightConfig.User == "" {
		errorLog.Fatalln("HUE_LIGHT_USER is not set")
	}
	if lightConfig.Host == "" {
		log.Println("HUE_LIGHT_HOST is not set, auto selecting host")
	}

	// Get information and display
	info, err := light.Info(lightConfig)
	if err != nil {
		errorLog.Fatalln(err)
	}
	log.Println(info)
}
