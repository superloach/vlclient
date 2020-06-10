package vlclient

import (
	"net/url"
)

func (v *VLClient) PlRandom() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_random")

	return s, v.getJSON("/requests/status", args, &s)
}
