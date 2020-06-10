package vlclient

type Stats struct {
	DecodedAudio        int     `json:"decodedaudio"`
	InputBitrate        float64 `json:"inputbitrate"`
	LostABuffers        int     `json:"lostabuffers"`
	ReadBytes           int     `json:"readbytes"`
	AverageDemuxBitware float64 `json:"averagedemuxbitrate"`
	DemuxReadBytes      int     `json:"demuxreadbytes"`
	PlayedABuffers      int     `json:"playedabuffers"`
	DemuxCorrupted      int     `json:"demuxcorrupted"`
	DisplayedPictures   int     `json:"displayedpictures"`
	LostPictures        int     `json:"lostpictures"`
	SentPackets         int     `json:"sentpackets"`
	DemuxDiscontinuity  int     `json:"demuxdiscontinuity"`
	DecodedVideo        int     `json:"decodedvideo"`
	SendBitrate         float64 `json:"sendbitrate"`
	SentBytes           int     `json:"sentbytes"`
	DemuxReadPackets    int     `json:"demuxreadpackets"`
	DemuxBitrate        float64 `json:"demuxbitrate"`
	AverageInputBitrate float64 `json:"averageinputbitrate"`
	ReadPackets         int     `json:"readpackets"`
}
