package telegrambotmodel

// Repository represents the repository information in the payload.
type Repository struct {
	FullName string `json:"full_name"`
	HTMLURL  string `json:"html_url"`
}

// PullRequest represents pull request information in the payload.
type PullRequest struct {
	Title string `json:"title"`
	User  User   `json:"user"`
	URL   string `json:"html_url"`
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
