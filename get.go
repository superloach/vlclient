package vlclient

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (v *VLClient) get(endpoint string, args url.Values) (io.ReadCloser, error) {
	u, err := url.Parse(v.base.String())
	if err != nil {
		return nil, err
	}

	u.Path = endpoint
	us := u.String()
	a := args.Encode()
	us += "?" + strings.ReplaceAll(a, "+", "%20")

	req, err := http.NewRequest("GET", us, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", v.passwd)

	resp, err := v.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (v *VLClient) getJSON(endpoint string, args url.Values, value interface{}) error {
	body, err := v.get(endpoint+".json", args)
	if err != nil {
		return err
	}
	defer body.Close()

	return json.NewDecoder(body).Decode(value)
}
