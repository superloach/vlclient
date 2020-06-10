package vlclient

import (
	"net/url"
)

func (v *VLClient) PlSDRemove(val string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_sd_remove")
	args.Set("val", val)

	return s, v.getJSON("/requests/status", args, &s)
}
