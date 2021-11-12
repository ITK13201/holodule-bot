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

func loadEnv() {
	err := godotenv.Load()
	if err == nil {
		log.Default().Print("Loading .env file...")
	}
}

func init() {
	loadEnv()

	var cfg Config
	cfg.DiscordChannelId = mustGetEnv("DISCORD_CHANNEL_ID")
	cfg.DiscordBotToken = mustGetEnv("DISCORD_BOT_TOKEN")
	Cfg = &cfg

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	JST = jst
}
