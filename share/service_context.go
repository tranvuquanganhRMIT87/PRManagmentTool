package share

import (
	"ngrok-go-quickstart/Components/chatGPT"
	githubs "ngrok-go-quickstart/Components/github"
	internal "ngrok-go-quickstart/Components/telegram"
)

type serviceContext struct {
	TelegramBot internal.TelegramBot
	GithubAPI   githubs.GithubApi
	OpenAI      chatGPT.CodeReviewAssistant
	Config      EnvConfig
}

func NewServiceContext(TelegramBot internal.TelegramBot, Config EnvConfig, GithubAPI githubs.GithubApi) *serviceContext {
	return &serviceContext{
		TelegramBot: TelegramBot,
		Config:      Config,
		GithubAPI:   GithubAPI,
	}
}

func (s *serviceContext) GetConfig() EnvConfig {
	return s.Config
}
func (s *serviceContext) GetTelegramBot() internal.TelegramBot {
	return s.TelegramBot
}
func (s *serviceContext) GetGithubAPI() githubs.GithubApi {
	return s.GithubAPI
}
func (s *serviceContext) GetOpenAI() chatGPT.CodeReviewAssistant {
	return s.OpenAI
}

type ServiceContext interface {
	GetConfig() EnvConfig
	GetTelegramBot() internal.TelegramBot
	GetGithubAPI() githubs.GithubApi
	GetOpenAI() chatGPT.CodeReviewAssistant
}
