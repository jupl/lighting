package light

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/hue"
	"github.com/heatxsink/go-hue/lights"
	"github.com/heatxsink/go-hue/portal"
	"github.com/lucasb-eyer/go-colorful"
)

// Config used in lighting API
type Config struct {
	Host string
	User string
	ID   int
}

// Source defines a light source
type Source struct {
	groups *groups.Groups
	group  groups.Group
}

// Info returns a string containing host information.
func Info(config Config) (string, error) {
	var buffer bytes.Buffer

	// Verify config
	config, err := config.verify()
	if err != nil {
		return "", err
	}

	// Iterate over groups
	grps, err := groups.New(config.Host, config.User).GetAllGroups()
	if err != nil {
		return "", errors.New("Failed to parse groups")
	}
	if len(grps) > 0 {
		buffer.WriteString(fmt.Sprintln(""))
		buffer.WriteString(fmt.Sprintln("--------- Groups"))
		buffer.WriteString(fmt.Sprintln(""))
	}
	for _, group := range grps {
		buffer.WriteString(fmt.Sprintln(group.String()))
	}

	return buffer.String(), nil
}

// Find a light source with a given host and user. If host is blank, a host is
// automatically selected. If a light source with an id is not found then
// an error is returned.
func Find(config Config) (Source, error) {
	// Verify config
	config, err := config.verify()
	if err != nil {
		return Source{}, err
	}

	// Attempt to get list of groups
	all := groups.New(config.Host, config.User)
	grps, err := all.GetAllGroups()
	if err != nil {
		return Source{}, errors.New("Failed to parse groups")
	}

	// Iterate over all groups and find a group with a matching ID
	for _, group := range grps {
		if group.ID == config.ID {
			return Source{groups: all, group: group}, nil
		}
	}

	// Failed to find a light source
	return Source{}, errors.New("Cannot find light group for given ID")
}

// TurnOff a light source.
func (source Source) TurnOff() ([]hue.ApiResponse, error) {
	if source.groups == nil {
		return nil, errors.New("Cannot set state on an invalid light")
	}
	return source.groups.SetGroupState(source.group.ID, lights.State{On: false})
}

// TurnOn light source to a given state.
func (source Source) TurnOn(color colorful.Color) ([]hue.ApiResponse, error) {
	if source.groups == nil {
		return nil, errors.New("Cannot set state on an invalid light")
	}
	hue, sat, val := color.Hsv()
	return source.groups.SetGroupState(source.group.ID, lights.State{
		On:  true,
		Hue: uint16(hue / 360 * math.MaxUint16),
		Sat: uint8(math.Max(1, math.Min(math.MaxUint8-1, sat*math.MaxUint8))),
		Bri: uint8(math.Max(1, math.Min(math.MaxUint8-1, val*math.MaxUint8))),
	})
}

// verify checks config and attempts to fix any issues
func (config Config) verify() (Config, error) {
	if config.Host != "" {
		return config, nil
	}

	// Find a portal to get a host
	portals, err := portal.GetPortal()
	if err != nil {
		return config, errors.New("Failed to parse portal")
	} else if len(portals) == 0 {
		return config, errors.New("No bridges available")
	}
	config.Host = portals[0].InternalIPAddress
	return config, nil
}
