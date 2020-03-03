package model

// Creator ...
type Creator struct {
	UID         string               `json:"uid"`      // identificador interno unificado
	Identity    string               `json:"identity"` // identificador único e memorável ("arroba")
	Description string               `json:"description"`
	About       string               `json:"about"` // descrição mais longa. TODO: aceitar markdown
	Links       SocialLinkCollection `json:"links"`
}
