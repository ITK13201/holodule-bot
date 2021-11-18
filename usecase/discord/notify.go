package discord

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ITK13201/holodule-bot/config"
	"github.com/ITK13201/holodule-bot/usecase/video"
)

type Author struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	IconUrl string `json:"icon_url"`
}

type Image struct {
	Url string `json:"url"`
}

type Embed struct {
	Author      Author `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       Image  `json:"image"`
}

func GetEmbed(video video.VideoWithDatetime) Embed {
	datetime := video.Datetime.Format(time.ANSIC)

	content := Embed{
		Author: Author{
			Name:    video.Distributor.Name,
			Url:     "",
			IconUrl: video.Distributor.IconUrl,
		},
		Title:       datetime,
		Description: video.Url,
		Image: Image{
			Url: video.ImageUrl,
		},
	}

	return content
}

type EmbedsContent struct {
	Embeds []Embed `json:"embeds"`
}

type OnlyTextContent struct {
	Content string `json:"content"`
}

func NotifyWithWebhook(contentJson []byte, url string) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(contentJson))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("[!] " + err.Error())
	} else {
		log.Info("[*] " + resp.Status)
	}

	log.Info(bytes.NewBuffer(contentJson).String())

	respTxt, _ := ioutil.ReadAll(resp.Body)
	log.Info(string(respTxt))
}

func NotifyWIthBot(contentJson []byte, channelId string) {
	cfg := *config.Cfg
	url := fmt.Sprintf("https://discordapp.com/api/channels/%s/messages", channelId)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(contentJson))
	req.Header.Set("Authorization", fmt.Sprintf("Bot %s", cfg.DiscordBotToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Error("[!] " + err.Error())
	} else {
		log.Info("[*] " + resp.Status)
	}

	log.Info(bytes.NewBuffer(contentJson).String())

	respTxt, _ := ioutil.ReadAll(resp.Body)
	log.Info(string(respTxt))
}
