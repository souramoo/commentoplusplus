package main

import (
	"time"
)

type domainOwner struct {
	OwnerHex string    `json:"ownerHex"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	AddDate  time.Time `json:"addDate"`
}

type domain struct {
	Domain                  string        `json:"domain"`
	Owners                  []domainOwner `json:"owners"`
	Name                    string        `json:"name"`
	CreationDate            time.Time     `json:"creationDate"`
	State                   string        `json:"state"`
	ImportedComments        bool          `json:"importedComments"`
	AutoSpamFilter          bool          `json:"autoSpamFilter"`
	RequireModeration       bool          `json:"requireModeration"`
	RequireIdentification   bool          `json:"requireIdentification"`
	ModerateAllAnonymous    bool          `json:"moderateAllAnonymous"`
	Moderators              []moderator   `json:"moderators"`
	EmailNotificationPolicy string        `json:"emailNotificationPolicy"`
	CommentoProvider        bool          `json:"commentoProvider"`
	GoogleProvider          bool          `json:"googleProvider"`
	TwitterProvider         bool          `json:"twitterProvider"`
	GithubProvider          bool          `json:"githubProvider"`
	GitlabProvider          bool          `json:"gitlabProvider"`
	SsoProvider             bool          `json:"ssoProvider"`
	SsoSecret               string        `json:"ssoSecret"`
	SsoUrl                  string        `json:"ssoUrl"`
	DefaultSortPolicy       string        `json:"defaultSortPolicy"`
}
