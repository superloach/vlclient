package vlclient

import (
	"net/url"
)

func (v *VLClient) PlForceResume() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_forceresume")

	return s, v.getJSON("/requests/status", args, &s)
}
