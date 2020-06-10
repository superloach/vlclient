package vlclient

import (
	"net/url"
)

func (v *VLClient) EnableEq(state bool) (*Status, error) {
	s := &Status{}

	val := "0"
	if state {
		val = "1"
	}

	args := make(url.Values)
	args.Set("command", "enableeq")
	args.Set("val", val)

	return s, v.getJSON("/requests/status", args, &s)
}
