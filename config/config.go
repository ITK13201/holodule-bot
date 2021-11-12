package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordChannelId string `json:"discord-channel-id"`
	DiscordBotToken  string `json:"discord-bot-token"`
	DiscordWebhookUrl string `json:"discord-webhook-url"`
}

var Cfg *Config
var JST *time.Location

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic(fmt.Sprintf("Error: Environment valiable with key of \"%s\" cannnot be read", key))
	}
	return value
}

func loadDevEnv() {
	err := godotenv.Load("dev.env")
	if err == nil {
		log.Default().Print("Loading dev.env file...")
	}
}

func init() {
	loadDevEnv()

	var cfg Config
	cfg.DiscordChannelId = mustGetEnv("DISCORD_CHANNEL_ID")
	cfg.DiscordBotToken = mustGetEnv("DISCORD_BOT_TOKEN")
	cfg.DiscordWebhookUrl = mustGetEnv("DISCORD_WEBHOOK_URL")
	Cfg = &cfg

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	JST = jst
}
