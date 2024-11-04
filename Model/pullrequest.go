package Model

var Payload struct {
	Action      string `json:"action"`
	PullRequest struct {
		URL   string `json:"html_url"`
		Title string `json:"title"`
		User  struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"pull_request"`
	Repository struct {
		FullName string `json:"full_name"`
		HTMLURL  string `json:"html_url"`
	} `json:"repository"`
}
