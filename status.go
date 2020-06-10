package vlclient

type Status struct {
	SubtitleDelay float64           `json:"subtitledelay"`
	Random        bool              `json:"random"`
	APIVersion    int               `json:"apiversion"`
	Position      float64           `json:"position"`
	VideoEffects  VideoEffects      `json:"videoeffects"`
	Version       string            `json:"version"`
	AudioFilters  map[string]string `json:"audiofilters"`
	Time          int               `json:"time"`
	Rate          float64           `json:"rate"`
	State         string            `json:"state"`
	CurrentPLID   int               `json:"currentplid"`
	Stats         Stats             `json:"stats"`
	Information   Information       `json:"information"`
	Volume        float64           `json:"volume"`
	Equalizer     []interface{}     `json:"equalizer"`
	Fullscreen    int               `json:"fullscreen"`
	AudioDelay    float64           `json:"audiodelay"`
	Length        int               `json:"length"`
	Repeat        bool              `json:"repeat"`
	Loop          bool              `json:"loop"`
}

func (v *VLClient) Status() (*Status, error) {
	s := &Status{}

	err := v.getJSON("/requests/status", nil, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
