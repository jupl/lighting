package main

import (
	"github.com/joho/godotenv"
	"github.com/jupl/lighting/env"
	"github.com/jupl/lighting/light"
	"github.com/jupl/lighting/logger"
	"github.com/lucasb-eyer/go-colorful"
	"time"
)

var redColor, _ = colorful.Hex("#FF4136")
var greenColor, _ = colorful.Hex("#2ECC40")

func main() {
	log := logger.New()
	errorLog := logger.NewError()
	infoLog := logger.NewInfo()

	// Attempt to read .env file
	if godotenv.Load() != nil {
		errorLog.Println("Cannot read .env file")
	}

	// Read light config
	lightConfig, err := env.LightConfig(true)
	if err != nil {
		errorLog.Fatalln(err)
	}
	if lightConfig.Host == "" {
		log.Println("HUE_LIGHT_HOST is not set, auto selecting host")
	}

	for {
		// Attempt to update light, displaying any errors
		if err := updateLight(lightConfig); err != nil {
			infoLog.Println(err)
		}

		// Wait for time to pass
		time.Sleep(time.Minute)
	}
}

func updateLight(config light.Config) error {
	// Find a light source
	source, err := light.Find(config)
	if err != nil {
		return err
	}

	// Process light
	color := redColor
	if time.Now().Minute()%2 == 0 {
		color = greenColor
	}
	_, err = source.TurnOn(color)
	return err
}
