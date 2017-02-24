package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/jupl/lighting/light"
)

// LightConfig reads light configuration. ID checking can be disabled using
// checkID. Any issues with parsing the config results in na error.
func LightConfig(checkID bool) (light.Config, error) {
	id, _ := strconv.Atoi(os.Getenv("HUE_LIGHT_ID"))
	config := light.Config{
		Host: os.Getenv("HUE_LIGHT_HOST"),
		User: os.Getenv("HUE_LIGHT_USER"),
		ID:   id,
	}
	if config.User == "" {
		return config, errors.New("HUE_LIGHT_USER is not set")
	} else if checkID && config.ID == 0 {
		return config, errors.New("HUE_LIGHT_ID is not set")
	}
	return config, nil
}
