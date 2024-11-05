package share

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type EnvConfig interface {
	GetHttpAddr() string
	GetTelegramToken() string
	GetChatID() int64
	GetThreadID() int
	InitConfig()
}

type env struct{}

func (env *env) GetHttpAddr() string {
	return os.Getenv("PORT")
}
func (env *env) GetTelegramToken() string {
	return os.Getenv("TELE_BOT_TOKEN")
}
func (env *env) GetChatID() int64 {
	chatID, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("Invalid CHAT_ID: %v", err)
	}
	return chatID
}
func (env *env) GetThreadID() int {
	threadID, err := strconv.Atoi(os.Getenv("THREAD_ID"))
	if err != nil {
		log.Fatalf("Invalid THREAD_ID: %v", err)
	}
	return threadID
}

func (env *env) InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewEnvConfig() EnvConfig {
	return &env{}
}
