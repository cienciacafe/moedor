package model

import (
	"fmt"

	"ciencia.cafe/moedor/util"
)

type (
	// PostType representa o tipo do post ("video", "podcast" ou "text")
	PostType string
	// SourceType representa a origem do arquivo de mídia ("file" ou "youtube")
	SourceType string
)

// Tipos de post
const (
	VideoPost   PostType = "video"
	PodcastPost PostType = "podcast"
	TextPost    PostType = "text"
)

// Tipos de origem
const (
	File    SourceType = "file"
	Youtube SourceType = "youtube"
)

// PostMedia ...
type PostMedia struct {
	Duration int        `json:"duration"`
	FileSize int        `json:"fileSize"`
	FileType string     `json:"fileType"`
	URL      string     `json:"url"`
	Source   SourceType `json:"source"`
}

// Post ...
type Post struct {
	Timestamp int       `json:"timestamp"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Summary   string    `json:"summary"`
	GUID      string    `json:"guid"`
	ImageURL  string    `json:"imageURL"`
	Tags      []string  `json:"tags"`
	Type      PostType  `json:"type"`
	Media     PostMedia `json:"media"`
}

// Checksum cálcula um identificador para o post baseado no
// title, summary, imageURL, media.duration, media.url e media.fileSize
func (p Post) Checksum() string {
	return util.Hash(fmt.Sprintf(
		"%s;%s;%s;%d;%s;%d",
		p.Title,
		p.Summary,
		p.ImageURL,
		p.Media.Duration,
		p.Media.URL,
		p.Media.FileSize,
	))
}
