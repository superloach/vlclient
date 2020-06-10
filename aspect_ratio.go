package vlclient

import (
	"net/url"
)

func (v *VLClient) AspectRatio(val string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "aspectratio")
	args.Set("val", val)

	return s, v.getJSON("/requests/status", args, &s)
}
