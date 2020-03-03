package model

// ChannelType ...
type ChannelType string

// Tipo de "canal" (não pensei em um nome genérico melhor)
const (
	PodcastChannel ChannelType = "podcast"
	YoutubeChannel ChannelType = "youtube"
)

// Channel ...
type Channel struct {
	UID         string               `json:"uid"`      // identificador interno unificado
	Identity    string               `json:"identity"` // identificador único e memorável ("arroba")
	Title       string               `json:"title"`
	AvatarURL   string               `json:"imageURL"`
	Description string               `json:"description"`
	About       string               `json:"about"` // descrição mais longa. TODO: aceitar markdown
	Type        ChannelType          `json:"type"`
	FeedURL     string               `json:"feedURL"`
	Links       SocialLinkCollection `json:"links"`

	// IDEA: salvar lista de canais recomendados?
	// ex: https://www.youtube.com/user/CanalPeixeBabel/channels
}
