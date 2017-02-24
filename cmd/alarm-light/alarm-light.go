package main

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/jupl/lighting/env"
	"github.com/jupl/lighting/light"
	"github.com/jupl/lighting/logger"
	"github.com/lucasb-eyer/go-colorful"
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
	config, err := env.LightConfig(true)
	if err != nil {
		errorLog.Fatalln(err)
	}
	if config.Host == "" {
		log.Println("HUE_LIGHT_HOST is not set, auto selecting host")
	}

	// Attempt to update light per interval, displaying any errors
	for {
		if err := updateLight(config); err != nil {
			infoLog.Println(err)
		}
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
