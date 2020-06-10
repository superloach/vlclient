package vlclient

import (
	"errors"
	"net/url"
)

func (v *VLClient) PlPlay(ids ...string) (*Status, error) {
	s := &Status{}

	args := make(url.Values)
	args.Set("command", "pl_play")

	if len(ids) > 0 {
		return nil, errors.New("pl_play takes 0 or 1 ids")
	}

	for _, id := range ids {
		args.Set("id", id)
	}

	return s, v.getJSON("/requests/status", args, &s)
}
