package main

import (
	"flag"
	"fmt"
	"github.com/jupl/lighting/env"
	"github.com/jupl/lighting/light"
	"os"
)

var host = ""
var user = ""

func init() {
	flag.StringVar(&host, "host", env.LightHost(), "Light API host")
	flag.StringVar(&user, "user", env.LightUser(), "Light API user")
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
	if user == "" {
		usage()
	}

	// Note if hostname is not provided
	if host == "" {
		fmt.Println("No host provided, automatically selecting a host")
	}

	// Get information and display
	info, err := light.Info(host, user)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(info)
}
