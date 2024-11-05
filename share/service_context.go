package share

import (
	internal "ngrok-go-quickstart/Components/telegram"
)

type serviceContext struct {
	TelegramBot internal.TelegramBot
	Config      EnvConfig
}

func NewServiceContext(TelegramBot internal.TelegramBot, Config EnvConfig) *serviceContext {
	return &serviceContext{
		TelegramBot: TelegramBot,
		Config:      Config,
	}
}

func (s *serviceContext) GetConfig() EnvConfig {
	return s.Config
}
func (s *serviceContext) GetTelegramBot() internal.TelegramBot {
	return s.TelegramBot
}

type ServiceContext interface {
	GetConfig() EnvConfig
	GetTelegramBot() internal.TelegramBot
}
