package transport

import (
	"context"
	telegrambotmodel "ngrok-go-quickstart/modules/telegram_bot/model"
)

type PushMessageHandler interface {
	Execute(ctx context.Context, payload telegrambotmodel.Payload) (bool, error)
}
type httpService struct {
	pushMessageHandler PushMessageHandler
}

func NewHttpService(pushMessageHandler PushMessageHandler) *httpService {
	return &httpService{
		pushMessageHandler: pushMessageHandler,
	}
}
