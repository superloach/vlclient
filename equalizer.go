package vlclient

import (
	"errors"
	"net/url"
	"strconv"
)

const (
	Band60Hz = iota
	Band170Hz
	Band310Hz
	Band600Hz
	Band1kHz
	Band3kHz
	Band6kHz
	Band12kHz
	Band14kHz
	Band16kHz
)

func (v *VLClient) Equalizer(band, val int) (*Status, error) {
	s := &Status{}

	if band < Band60Hz || band > Band16kHz {
		return nil, errors.New("invalid band")
	}

	if val < -20 || val > 20 {
		return nil, errors.New("gain must be >=-20 and <=20")
	}

	args := make(url.Values)
	args.Set("command", "preamp")
	args.Set("val", strconv.Itoa(val))

	return s, v.getJSON("/requests/status", args, &s)
}
