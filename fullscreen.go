package vlclient

import (
	"net/url"
)

func (v *VLClient) Fullscreen() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "fullscreen")

	return s, v.getJSON("/requests/status", args, &s)
}
