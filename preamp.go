package vlclient

import (
	"errors"
	"net/url"
	"strconv"
)

func (v *VLClient) Preamp(val int) (*Status, error) {
	s := &Status{}

	if val < -20 || val > 20 {
		return nil, errors.New("preamp must be >=-20 and <=20")
	}

	args := make(url.Values)
	args.Set("command", "preamp")
	args.Set("val", strconv.Itoa(val))

	return s, v.getJSON("/requests/status", args, &s)
}
