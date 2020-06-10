package vlclient

type Information struct {
	Chapters []interface{} `json:"chapters"`
	Titles   []interface{} `json:"titles"`
	Title    int           `json:"title"`
	Chapter  int           `json:"chapter"`
	Category Category      `json:"category"`
}
