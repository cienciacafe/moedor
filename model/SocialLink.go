package model

// LinkType ...
type LinkType int

// Tipos de link
const (
	Website LinkType = iota
	YouTube
	Twitter
	Facebook
	Instagram
	Lattes
)

// SocialLinkCollection ...
type SocialLinkCollection []SocialLink

// SocialLink ...
type SocialLink struct {
	RelatedEntity string   `json:"relatedUID"` // identificador da 'entidade' (canal ou criador) relacionada ao link
	URL           string   `json:"url"`
	Type          LinkType `json:"type"`
}
