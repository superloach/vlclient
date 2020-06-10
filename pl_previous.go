package vlclient

import (
	"net/url"
)

func (v *VLClient) PlPrevious() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_previous")

	return s, v.getJSON("/requests/status", args, &s)
}
