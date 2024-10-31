package Model

type payload struct {
	Id          string `json:"id"`
	Action      string `json:"action"`
	Order       int    `json:"number"`
	PullRequest struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		HTMLURL     string `json:"html_url"`
	} `json:"pull_request"`
	Repository struct {
		Name string `json:"name"`
	} `json:"repository"`
	Sender struct {
		Login string `json:"login"`
	} `json:"sender"`
}
