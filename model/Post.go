package model

type PostType string
type SourceType string

const (
	VideoPost   PostType = "video"
	PodcastPost PostType = "podcast"
	TextPost    PostType = "text"
)

const (
	File    SourceType = "file"
	Youtube SourceType = "youtube"
)

type MediaSource struct {
	URL      string     `json:"URL"`
	Type     SourceType `json:"type"`
	Size     int        `json:"size"`
	FileType string     `json:"fileType"`
}

type PostMedia struct {
	Duration int
	Source   MediaSource `json:"source"`
}

type Post struct {
	Timestamp int      `json:"timestamp"`
	Type      PostType `json:"type"`
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	Summary   string   `json:"summary"`
	GUID      string   `json:"guid"`
	ImageURL  string   `json:"imageURL"`
	Media     string   `json:"media"`
}
