package vlclient

import (
	"errors"
	"net/url"
	"strconv"
)

func (v *VLClient) Rate(val float64) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "rate")

	if val <= 0 {
		return nil, errors.New("rate must be > 0")
	}

	args.Set("val", strconv.FormatFloat(val, 'f', -1, 64))

	return s, v.getJSON("/requests/status", args, &s)
}
