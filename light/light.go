package light

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/hue"
	"github.com/heatxsink/go-hue/lights"
	"github.com/heatxsink/go-hue/portal"
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

type states struct {
	Off    lights.State
	Red    lights.State
	Green  lights.State
	Orange lights.State
}

var noSource = Source{}

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

	// Verify host
	host, err := verifyHost(config.Host)
	if err != nil {
		return "", err
	}

	// Iterate over groups
	allGroups, err := groups.New(host, config.User).GetAllGroups()
	if err != nil {
		return "", errors.New("Failed to parse groups")
	}
	if len(allGroups) > 0 {
		buffer.WriteString(fmt.Sprintln(""))
		buffer.WriteString(fmt.Sprintln("--------- Groups"))
		buffer.WriteString(fmt.Sprintln(""))
	}
	for _, group := range allGroups {
		buffer.WriteString(fmt.Sprintln(group.String()))
	}

	return buffer.String(), nil
}

// Find a light source with a given host and user. If host is blank, a host is
// automatically selected. If a light source with an id is not found then
// an error is returned.
func Find(config Config) (Source, error) {
	// Verify host
	host, err := verifyHost(config.Host)
	if err != nil {
		return noSource, err
	}

	// Attempt to get list of groups
	groups := groups.New(host, config.User)
	allGroups, err := groups.GetAllGroups()
	if err != nil {
		return noSource, errors.New("Failed to parse groups")
	}

	// Iterate over all groups and find a group with a matching ID
	light := noSource
	err = errors.New("Cannot find light group for given ID")
	for _, group := range allGroups {
		if group.ID != config.ID {
			continue
		}
		light = Source{groups: groups, group: group}
		err = nil
		break
	}
	return light, err
}

// TurnOff a light source.
func (light Source) TurnOff() ([]hue.ApiResponse, error) {
	if light.groups == nil {
		return nil, errors.New("Cannot set state on an invalid light")
	}
	return light.groups.SetGroupState(light.group.ID, lights.State{On: false})
}

// TurnOn light source to a given state.
func (light Source) TurnOn(color colorful.Color) ([]hue.ApiResponse, error) {
	if light.groups == nil {
		return nil, errors.New("Cannot set state on an invalid light")
	}
	hue, sat, bri := color.Hsv()
	return light.groups.SetGroupState(light.group.ID, lights.State{
		On:  true,
		Hue: uint16(hue / 360 * math.MaxUint16),
		Sat: uint8(math.Max(1, math.Min(math.MaxUint8-1, sat*math.MaxUint8))),
		Bri: uint8(math.Max(1, math.Min(math.MaxUint8-1, bri*math.MaxUint8))),
	})
}

func verifyHost(host string) (string, error) {
	if host == "" {
		portals, err := portal.GetPortal()
		if err != nil {
			return "", errors.New("Failed to parse portal")
		} else if len(portals) == 0 {
			return "", errors.New("No bridges available")
		}
		host = portals[0].InternalIPAddress
	}
	return host, nil
}
