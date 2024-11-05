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

	if payload.Action == "opened" {
		message := fmt.Sprintf(
			"🆕: New Pull Request in repository %s:\nTitle: %s\nBy: %s\nPR URL: %s\nRepository URL: %s",
			payload.Repository.FullName,
			payload.PullRequest.Title,
			payload.PullRequest.User.Login,
			payload.PullRequest.URL,
			payload.Repository.HTMLURL,
		)

		formatedMessage := components.NewBotMessage(chatId, &threadId, message)

		if err := p.action.SendMessage(formatedMessage); err != nil {
			return false, err
		}
	}

	return true, nil

}
