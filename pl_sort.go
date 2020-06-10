package vlclient

import (
	"net/url"
	"strconv"
)

const (
	SortID          = 0
	SortName        = 1
	SortAuthor      = 3
	SortRandom      = 5
	SortTrackNumber = 7
)

func (v *VLClient) PlSort(mode int, rev bool) (*Status, error) {
	s := &Status{}

	id := "0"
	if rev {
		id = "1"
	}

	args := make(url.Values)
	args.Set("command", "pl_sort")
	args.Set("id", id)
	args.Set("val", strconv.Itoa(mode))

	return s, v.getJSON("/requests/status", args, &s)
}
