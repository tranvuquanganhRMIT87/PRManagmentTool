package components

import (
	"fmt"
	tgbotapi "github.com/nghiatrann0502/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

type TelegramBot interface {
	Connect() error
	SendMessage(message Message) error
}

type Telegram struct {
	Token string
	Bot   *tgbotapi.BotAPI
}

type Message struct {
	ChatID   int64
	ThreadId *int
	Text     string
}

func NewTelegram(token string) *Telegram {
	return &Telegram{
		Token: token,
	}
}

func NewBotMessage(chatId int64, threadId *int, text string) Message {
	return Message{
		ChatID:   chatId,
		ThreadId: threadId,
		Text:     text,
	}
}

func (t *Telegram) Connect() error {
	bot, err := tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		return err
	}

	bot.Debug = true
	t.Bot = bot
	logrus.Info("Telegram bot connected")

	return nil
}

func (t *Telegram) SendMessage(message Message) error {
	if t.Bot == nil {
		return fmt.Errorf("telegram bot is not initialized")
	}

	msg := tgbotapi.NewMessage(message.ChatID, message.Text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	if message.ThreadId != nil {
		msg.ReplyToMessageID = *message.ThreadId
	}

	_, err := t.Bot.Send(msg)
	if err != nil {
		return err
	}
	logrus.Info("Message sent successfully to chat ID:", message.ChatID)
	return nil
}
