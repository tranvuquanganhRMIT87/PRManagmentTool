package telegrambotmodel

import "strings"

// Repository represents the repository information in the payload.
type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	HTMLURL  string `json:"html_url"`
}

// PullRequest represents pull request information in the payload.
type PullRequest struct {
	Title  string `json:"title"`
	User   User   `json:"user"`
	URL    string `json:"html_url"`
	Number int    `json:"number"`
	Head   Head   `json:"head"`
	Base   Base   `json:"base"`
}

type Head struct {
	Ref string `json:"ref"`
}
type Base struct {
	Ref string `json:"ref"`
}

// User represents the user information in the pull request.
type User struct {
	Login string `json:"login"`
}

// Commit represents individual commits in the payload.
type Commit struct {
	Message string `json:"message"`
}

type Payload struct {
	Action      string      `json:"action"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pull_request"`
	Commits     []Commit    `json:"commits"`
}

func (p *Payload) GetOwner() string {
	parts := strings.Split(p.Repository.FullName, "/")
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}
