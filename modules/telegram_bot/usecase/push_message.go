package usecase

import (
	"context"
	"fmt"
	"github.com/google/go-github/v66/github"
	"github.com/sirupsen/logrus"
	"log"
	"ngrok-go-quickstart/Components/chatGPT"
	githubs "ngrok-go-quickstart/Components/github"
	components "ngrok-go-quickstart/Components/telegram"
	telegrambotmodel "ngrok-go-quickstart/modules/telegram_bot/model"
	"ngrok-go-quickstart/share"
)

type pushMessageUseCase struct {
	action    components.TelegramBot
	config    share.EnvConfig
	githubAPI githubs.GithubApi
	openai    chatGPT.CodeReviewAssistant
}

func NewPushMessageUseCase(action components.TelegramBot, config share.EnvConfig, githubAPI githubs.GithubApi, openai chatGPT.CodeReviewAssistant) *pushMessageUseCase {
	return &pushMessageUseCase{
		action,
		config,
		githubAPI,
		openai,
	}
}

func (p *pushMessageUseCase) Execute(ctx context.Context, payload telegrambotmodel.Payload) (bool, error) {

	for _, commit := range payload.Commits {
		if len(commit.Message) >= 5 && (commit.Message[:5] == "Merge" || commit.Message[:7] == "Merged ") {
			return false, nil
		}
	}

	chatId := p.config.GetChatID()
	threadId := p.config.GetThreadID()

	var message string
	//pr, _ := share.ExtractPRNumber(payload.PullRequest.URL)

	switch payload.Action {
	case "opened":
		message = fmt.Sprintf(
			"**🔥 New Pull Request in repository [%s](%s):**\n\n"+
				"📝 **Title**\n"+
				"> %s\n\n"+
				"👤 **By:** %s\n"+
				"🌐 **PR URL:** [%s](%s)",
			payload.Repository.FullName,
			payload.Repository.HTMLURL,
			payload.PullRequest.Title,
			payload.PullRequest.User.Login,
			payload.PullRequest.URL,
			payload.PullRequest.URL,
		)
	case "closed":

		message = fmt.Sprintf(
			"**✅ Pull Request #%s Merged!**\n\n"+
				"🎉 **Congratulations!** The pull request has been successfully merged!\n\n"+
				"👤 **Merged By:** %s\n\n"+
				"🚀 **Happy coding!**",
			payload.PullRequest.Number,
			payload.PullRequest.User.Login,
		)

	}

	if message != "" {
		formattedMessage := components.NewBotMessage(chatId, &threadId, message)

		if err := p.action.SendMessage(formattedMessage); err != nil {
			return false, err
		}
	}

	owner := payload.GetOwner()
	repo := payload.Repository.FullName
	prNumber := payload.PullRequest.Number
	head := payload.PullRequest.Head.Ref
	base := payload.PullRequest.Base.Ref

	//files, err := p.githubAPI.ListPullRequestFiles(ctx, owner, repo, prNumber)
	//if err != nil {
	//	log.Fatalf("Failed to list PR files: %v", err)
	//}

	commitFiles, err := p.githubAPI.GetBranchDiff(ctx, owner, repo, base, head)
	if err != nil {
		log.Fatalf("Failed to get branch diff: %v", err)
	}

	var patchesList []*github.CommitFile
	for _, file := range commitFiles {
		if len(share.Encode(file.GetPatch())) <= share.MaxTokens/2 {
			patchesList = append(patchesList, file)
		} else {
			logrus.Printf("Skipping file %s due to large patch size", file.GetFilename())
		}
	}

	for _, file := range patchesList {
		if suggestionText, err := p.openai.GetOpenAiSuggestions(ctx, file.GetPatch()); err == nil {
			firstChangedLine := share.ExtractFirstChangedLine(file.GetPatch())
			lastCommit, err := p.githubAPI.GetLastCommit(ctx, owner, repo, prNumber)
			if err != nil {
				return false, err
			}
			path := file.GetFilename()
			comment := &github.PullRequestComment{
				Body:     github.String(fmt.Sprintf("[ChatGPTReviewer]\n%s", suggestionText)),
				Path:     &path,
				CommitID: &lastCommit,
				Line:     &firstChangedLine,
			}
			if err := p.githubAPI.CreateComment(ctx, owner, repo, prNumber, comment); err != nil {
				return false, err
			}
		}
	}

	return true, nil
}
