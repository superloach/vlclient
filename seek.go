package vlclient

import (
	"net/url"
)

func (v *VLClient) Seek(val string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "seek")
	args.Set("val", val)

	return s, v.getJSON("/requests/status", args, &s)
}
