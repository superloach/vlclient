package vlclient

import (
	"net/url"
)

func (v *VLClient) PlDelete(id string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_delete")
	args.Set("id", id)

	return s, v.getJSON("/requests/status", args, &s)
}
