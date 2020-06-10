package vlclient

import (
	"net/url"
)

func (v *VLClient) PlEmpty() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_empty")

	return s, v.getJSON("/requests/status", args, &s)
}
