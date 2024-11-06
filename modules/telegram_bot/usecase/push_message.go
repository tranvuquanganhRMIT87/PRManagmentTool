package usecase

import (
	"context"
	"fmt"
	components "ngrok-go-quickstart/Components/telegram"
	telegrambotmodel "ngrok-go-quickstart/modules/telegram_bot/model"
	"ngrok-go-quickstart/share"
)

type pushMessageUseCase struct {
	action components.TelegramBot
	config share.EnvConfig
}

func NewPushMessageUseCase(action components.TelegramBot, config share.EnvConfig) *pushMessageUseCase {
	return &pushMessageUseCase{
		action,
		config,
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
	pr, _ := share.ExtractPRNumber(payload.PullRequest.URL)

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
			pr,
			payload.PullRequest.User.Login,
		)

	}

	formattedMessage := components.NewBotMessage(chatId, &threadId, message)

	if err := p.action.SendMessage(formattedMessage); err != nil {
		return false, err
	}

	return true, nil
}
