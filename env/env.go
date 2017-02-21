package env

import (
	"github.com/jupl/lighting/light"
	"os"
	"strconv"
)

// LightConfig reads light configuration.
func LightConfig() light.Config {
	id, _ := strconv.Atoi(os.Getenv("HUE_LIGHT_ID"))
	return light.Config{
		Host: os.Getenv("HUE_LIGHT_HOST"),
		User: os.Getenv("HUE_LIGHT_USER"),
		ID:   id,
	}
}
