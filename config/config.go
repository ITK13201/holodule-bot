package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUrl                 string `json:"database-url" envconfig:"DATABASE_URL" required:"true"`
	DiscordBotToken             string `json:"discord-bot-token" envconfig:"DISCORD_BOT_TOKEN" required:"true"`
	DiscordChannelIdDaily       string `json:"discord-channel-id-daily" envconfig:"DISCORD_CHANNEL_ID_DAILY" required:"true"`
	DiscordChannelIdComingSoon  string `json:"discord-channel-id-coming-soon" envconfig:"DISCORD_CHANNEL_ID_COMING_SOON" required:"true"`
	DiscordWebhookUrlDaily      string `json:"discord-webhook-url-daily" envconfig:"DISCORD_WEBHOOK_URL_DAILY" required:"true"`
	DiscordWebhookUrlComingSoon string `json:"discord-webhook-url-coming-soon" envconfig:"DISCORD_WEBHOOK_URL_COMING_SOON" required:"true"`
	YoutubeDataAPIKey           string `json:"youtube-data-api-key" envconfig:"YOUTUBE_DATA_API_KEY" required:"true"`
}

var Cfg *Config
var JST *time.Location
var UTC *time.Location

func loadDevEnv() {
	err := godotenv.Load("dev.env")
	if err == nil {
		log.Default().Print("Loading dev.env file...")
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	return &cfg, err
}

func init() {
	loadDevEnv()

	cfg, err := LoadConfig()
	Cfg = cfg

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	JST = jst

	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	UTC = utc
}
