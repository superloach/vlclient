package vlclient

import "net/url"

func (v *VLClient) InEnqueue(input string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "in_enqueue")
	args.Set("input", input)

	return s, v.getJSON("/requests/status", args, &s)
}
