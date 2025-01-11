package githubs

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
	GetBranchDiff(ctx context.Context, owner, repo, base, head string) ([]*github.CommitFile, error)
	GetLastCommit(ctx context.Context, owner, repo string, number int) (string, error)
	CreateComment(ctx context.Context, owner, repo string, number int, comment *github.PullRequestComment) error
}

type githubC struct {
	Token  string
	Client *github.Client
}

func NewGithub(token string) *githubC {
	return &githubC{
		Token: token,
	}
}

func (g *githubC) Connect(ctx context.Context, httpClient *http.Client) error {
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

func (g *githubC) ListPullRequestFiles(ctx context.Context, owner, repo string, prNumber int) ([]*github.CommitFile, error) {
	files, _, err := g.Client.PullRequests.ListFiles(ctx, owner, repo, prNumber, nil)
	if err != nil {
		logrus.Errorf("Error fetching pull request files: %v", err)
		return nil, err
	}

	logrus.Infof("Fetched %d files from PR #%d in %s/%s", len(files), prNumber, owner, repo)
	return files, nil
}

func (g *githubC) GetBranchDiff(ctx context.Context, owner, repo, base, head string) ([]*github.CommitFile, error) {
	files, _, err := g.Client.Repositories.CompareCommits(ctx, owner, repo, base, head, nil)
	if err != nil {
		return nil, err
	}

	return files.Files, nil
}

func (g *githubC) GetLastCommit(ctx context.Context, owner, repo string, number int) (string, error) {
	commits, _, err := g.Client.PullRequests.ListCommits(ctx, owner, repo, number, nil)
	if err != nil {
		logrus.Errorf("Error fetching last commit: %v", err)
		return "", err
	}

	return commits[len(commits)-1].GetSHA(), nil
}

func (g *githubC) CreateComment(ctx context.Context, owner, repo string, number int, comment *github.PullRequestComment) error {
	_, _, err := g.Client.PullRequests.CreateComment(ctx, owner, repo, number, comment)
	if err != nil {
		logrus.Errorf("Error creating comment: %v", err)
		return err
	}

	logrus.Infof("Created comment on PR #%d in %s/%s", number, owner, repo)
	return nil
}
