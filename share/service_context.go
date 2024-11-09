package share

import (
	githubs "ngrok-go-quickstart/Components/github"
	internal "ngrok-go-quickstart/Components/telegram"
)

type serviceContext struct {
	TelegramBot internal.TelegramBot
	GithubAPI   githubs.GithubApi
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

type ServiceContext interface {
	GetConfig() EnvConfig
	GetTelegramBot() internal.TelegramBot
	GetGithubAPI() githubs.GithubApi
}
