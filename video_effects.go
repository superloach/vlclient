package vlclient

type VideoEffects struct {
	Hue        float64 `json:"hue"`
	Brightness float64 `json:"brightness"`
	Contrast   float64 `json:"contrast"`
	Saturation float64 `json:"saturation"`
	Gamma      float64 `json:"gamma"`
}
