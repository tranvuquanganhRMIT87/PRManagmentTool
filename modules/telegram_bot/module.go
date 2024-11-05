package telegram_bot

import (
	"net/http"
	"ngrok-go-quickstart/modules/telegram_bot/infras/transport"
	"ngrok-go-quickstart/modules/telegram_bot/usecase"
	"ngrok-go-quickstart/share"
)

func SetupTelegramBotService(sctx share.ServiceContext, mux *http.ServeMux) {

	pushMsgUC := usecase.NewPushMessageUseCase(sctx.GetTelegramBot(), sctx.GetConfig())

	pushMsgHdl := transport.NewHttpService(pushMsgUC)

	mux.HandleFunc("/webhook", pushMsgHdl.GithubWebhookHandler)

}
