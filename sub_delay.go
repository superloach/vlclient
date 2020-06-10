package vlclient

import (
	"net/url"
	"strconv"
)

func (v *VLClient) SubDelay(val float64) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "subdelay")
	args.Set("val", strconv.FormatFloat(val, 'f', -1, 64))

	return s, v.getJSON("/requests/status", args, &s)
}
