package vlclient

import (
	"net/http"
	"net/url"
)

type VLClient struct {
	client *http.Client

	base   *url.URL
	passwd string
}

func NewVLClient(base, passwd string) (*VLClient, error) {
	v := &VLClient{}

	v.client = &http.Client{}

	b, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	v.base = b

	v.passwd = passwd

	return v, nil
}
