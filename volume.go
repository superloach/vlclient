package vlclient

import (
	"net/url"
)

func (v *VLClient) Volume(val string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "volume")
	args.Set("val", val)

	return s, v.getJSON("/requests/status", args, &s)
}
