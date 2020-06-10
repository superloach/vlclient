package vlclient

import (
	"net/url"
)

func (v *VLClient) PlStop() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_stop")

	return s, v.getJSON("/requests/status", args, &s)
}
