package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	log "github.com/sirupsen/logrus"
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

type EmbedContent struct {
	Author      Author `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       Image  `json:"image"`
}

func getEmbedContent(video video.VideoWithDatetime) EmbedContent {
	datetime := video.Datetime.Format(time.ANSIC)

	content := EmbedContent{
		Author: Author{
			Name:   video.Distributor.Name,
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

type Content struct {
	Embeds []EmbedContent `json:"embeds"`
}

type OnlyTextContent struct {
	Content string `json:"content"`
}

func Notify(videos []video.VideoWithDatetime) {
	cfg := *config.Cfg
	url := fmt.Sprintf("https://discordapp.com/api/channels/%s/messages", cfg.DiscordChannelId)


	//=================
	// single text
	//=================
	now := time.Now().In(config.JST)
	content := OnlyTextContent{
		Content: fmt.Sprintf("**[INFO] %s**: Holodule updated", now.Format(time.ANSIC)),
	}
	contentJson, _ := json.Marshal(content)

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

	//====================
	// notify videos
	//====================
	allEmbedContents := []EmbedContent{}
	for i := 0; i < len(videos); i++ {
		allEmbedContents = append(allEmbedContents, getEmbedContent(videos[i]))
	}

	limit := len(allEmbedContents) / 10 + 1
	for i := 0; i < limit; i++ {
		var x = i*10
		var embedContents []EmbedContent
		if i == limit - 1 {
			embedContents = allEmbedContents[x:]
		} else {
			embedContents = allEmbedContents[x:x+10]
		}


		content := Content{
			Embeds: embedContents,
		}
		contentJson, _ := json.Marshal(content)

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
}
