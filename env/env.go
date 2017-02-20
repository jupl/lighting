package env

import (
	"os"
	"strconv"
)

// LightID provides a light group ID from the environment variable
// HUE_LIGHT_ID. If the environment variable cannot be parsed, the ID returned
// is 0.
func LightID() int {
	lightID, err := strconv.Atoi(os.Getenv("HUE_LIGHT_ID"))
	if err == nil {
		lightID = 0
	}
	return lightID
}

// LightUser gets the username for the light API from the environment
// variable HUE_LIGHT_USER.
func LightUser() string {
	return os.Getenv("HUE_LIGHT_USER")
}

// LightHost gets the hostname for the light API from the environment
// variable HUE_LIGHT_HOST.
func LightHost() string {
	return os.Getenv("HUE_LIGHT_HOST")
}
