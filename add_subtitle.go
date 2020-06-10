package vlclient

import "net/url"

func (v *VLClient) AddSubtitle(val string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "addsubtitle")
	args.Set("val", val)

	return s, v.getJSON("/requests/status", args, &s)
}
