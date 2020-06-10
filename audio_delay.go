package vlclient

import (
	"net/url"
	"strconv"
)

func (v *VLClient) AudioDelay(val float64) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "audiodelay")
	args.Set("val", strconv.FormatFloat(val, 'f', -1, 64))

	return s, v.getJSON("/requests/status", args, &s)
}
