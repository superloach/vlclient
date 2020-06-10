package vlclient

import (
	"net/url"
)

func (v *VLClient) PlLoop() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_loop")

	return s, v.getJSON("/requests/status", args, &s)
}
