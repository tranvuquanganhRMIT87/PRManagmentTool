package telegrambotmodel

import "errors"

var (
	ErrMissingGitHubToken     = errors.New("missing GitHub token")
	ErrMissingOpenAIToken     = errors.New("missing OpenAI API key")
	ErrNoPullRequestInContext = errors.New("no pull request in context")
	ErrNoChangedFilesInPR     = errors.New("no changed files in pull request")
)
