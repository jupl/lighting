package light

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/hue"
	"github.com/heatxsink/go-hue/lights"
	"github.com/heatxsink/go-hue/portal"
)

type states struct {
	Off    lights.State
	Red    lights.State
	Green  lights.State
	Orange lights.State
}

var noSource = Source{}

// Source defines a light source
type Source struct {
	groups *groups.Groups
	group  groups.Group
}

// States of lights
var States = states{
	Off: lights.State{On: false},
	Red: lights.State{
		On:  true,
		Hue: 65535,
		Sat: 254,
		Bri: 254,
	},
	Orange: lights.State{
		On:  true,
		Hue: 5460,
		Sat: 254,
		Bri: 254,
	},
	Green: lights.State{
		On:  true,
		Hue: 25500,
		Sat: 254,
		Bri: 254,
	},
}

// Info returns a string containing host information.
func Info(host string, user string) (string, error) {
	var buffer bytes.Buffer

	// Verify host
	host, err := verifyHost(host)
	if err != nil {
		return "", err
	}

	// Iterate over groups
	allGroups, err := groups.New(host, user).GetAllGroups()
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
func Find(host string, user string, id int) (Source, error) {
	// Verify host
	host, err := verifyHost(host)
	if err != nil {
		return noSource, err
	}

	// Attempt to get list of groups
	groups := groups.New(host, user)
	allGroups, err := groups.GetAllGroups()
	if err != nil {
		return noSource, errors.New("Failed to parse groups")
	}

	// Iterate over all groups and find a group with a matching ID
	light := noSource
	err = errors.New("Cannot find light group for given ID")
	for _, group := range allGroups {
		if group.ID != id {
			continue
		}
		light = Source{groups: groups, group: group}
		err = nil
		break
	}
	return light, err
}

// SetState for a light source to a given light state.
func (light Source) SetState(state lights.State) ([]hue.ApiResponse, error) {
	if light.groups == nil {
		return nil, errors.New("Cannot set state on an invalid light")
	}
	return light.groups.SetGroupState(light.group.ID, state)
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