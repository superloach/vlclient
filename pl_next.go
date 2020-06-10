package vlclient

import (
	"net/url"
)

func (v *VLClient) PlNext() (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_next")

	return s, v.getJSON("/requests/status", args, &s)
}
