package vlclient

import (
	"net/url"
)

func (v *VLClient) PlForcePause() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_forcepause")

	return s, v.getJSON("/requests/status", args, &s)
}
