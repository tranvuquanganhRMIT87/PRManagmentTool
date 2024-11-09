package components

import (
	"context"
	"github.com/google/go-github/v66/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type GithubApi interface {
	Connect(ctx context.Context, httpClient *http.Client) error
	ListPullRequestFiles(ctx context.Context, owner, repo string, prNumber int) ([]*github.CommitFile, error)
}

type Github struct {
	Token  string
	Client *github.Client
}

func (g *Github) Connect(ctx context.Context, httpClient *http.Client) error {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	// Setup token-based authentication
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: g.Token})
	tc := oauth2.NewClient(ctx, ts)
	g.Client = github.NewClient(tc)

	logrus.Info("Connected to GitHub API")
	return nil
}

func (g *Github) ListPullRequestFiles(ctx context.Context, owner, repo string, prNumber int) ([]*github.CommitFile, error) {
	files, _, err := g.Client.PullRequests.ListFiles(ctx, owner, repo, prNumber, nil)
	if err != nil {
		logrus.Errorf("Error fetching pull request files: %v", err)
		return nil, err
	}

	logrus.Infof("Fetched %d files from PR #%d in %s/%s", len(files), prNumber, owner, repo)
	return files, nil
}
