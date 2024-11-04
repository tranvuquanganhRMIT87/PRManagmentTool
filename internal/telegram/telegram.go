package internal

import tgbotapi "github.com/nghiatrann0502/telegram-bot-api"

type Telegram struct {
	Token string
	Bot   *tgbotapi.BotAPI
}

type Message struct {
	ChatID   int64
	ThreadId *int
	Text     string
}

// constructor
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

	return nil
}

func (t *Telegram) SendMessage(message Message) error {
	msg := tgbotapi.NewMessage(message.ChatID, message.Text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	// ?????
	if message.ThreadId != nil {
		msg.ReplyToMessageID = *message.ThreadId
	}
	_, err := t.Bot.Send(msg)

	if err != nil {
		return err
	}

	return nil
}
