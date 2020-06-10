package vlclient

import "net/url"

func (v *VLClient) InPlay(input string, option ...string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "in_play")
	args.Set("input", input)

	for _, opt := range option {
		args.Add("option", opt)
	}

	return s, v.getJSON("/requests/status", args, &s)
}
