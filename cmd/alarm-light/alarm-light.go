package main

import (
	"flag"
	"fmt"
	"github.com/jupl/lighting/env"
	"github.com/jupl/lighting/light"
	"github.com/lucasb-eyer/go-colorful"
	"os"
	"time"
)

var host = ""
var user = ""
var id = 0

var redColor, _ = colorful.Hex("#FF4136")
var greenColor, _ = colorful.Hex("#2ECC40")

func init() {
	flag.StringVar(&host, "host", env.LightHost(), "Light API host")
	flag.StringVar(&user, "user", env.LightUser(), "Light API user")
	flag.IntVar(&id, "id", env.LightID(), "Light group ID")
	flag.Parse()
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: alarm-light -host=[host] -user=[user] -id=[id]")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// Verify that data is available
	if user == "" || id == 0 {
		usage()
	}

	// Note if hostname is not provided
	if host == "" {
		fmt.Println("No host provided, automatically selecting a host")
	}

	for {
		// Attempt to update light, displaying any errors
		if err := updateLight(host, user, id); err != nil {
			fmt.Printf("[%s] %s\n", time.Now().Format(time.RFC822), err)
		}

		// Wait for time to pass
		time.Sleep(time.Minute)
	}
}

func updateLight(host string, user string, id int) error {
	// Find a light source
	source, err := light.Find(host, user, id)
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
