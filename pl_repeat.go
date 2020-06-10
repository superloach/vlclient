package vlclient

import (
	"net/url"
)

func (v *VLClient) PlRepeat() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_repeat")

	return s, v.getJSON("/requests/status", args, &s)
}
