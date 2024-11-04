package Model

var Payload struct {
	Action  string `json:"action"`
	Commits []struct {
		Message string `json:"message"`
	} `json:"commits"`
	PullRequest struct {
		URL   string `json:"html_url"`
		Title string `json:"title"`
		User  struct {
			Login string `json:"login"`
		} `json:"user"`
	} `json:"pull_request"`
	Events     string `json:"events"`
	Repository struct {
		FullName string `json:"full_name"`
		HTMLURL  string `json:"html_url"`
	} `json:"repository"`
}
