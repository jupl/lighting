package main

import (
	"flag"
	"fmt"
	"github.com/jupl/lighting/env"
	"github.com/jupl/lighting/light"
	"os"
)

var lightConfig = light.Config{}

func init() {
	flag.StringVar(&lightConfig.Host, "host", env.LightHost(), "Light API host")
	flag.StringVar(&lightConfig.User, "user", env.LightUser(), "Light API user")
	flag.Parse()
	flag.Usage = usage
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: lights -host=[host] -user=[user]")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	// Verify that data is available
	if lightConfig.User == "" {
		usage()
	}

	// Note if hostname is not provided
	if lightConfig.Host == "" {
		fmt.Println("No host provided, automatically selecting a host")
	}

	// Get information and display
	info, err := light.Info(lightConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(info)
}
