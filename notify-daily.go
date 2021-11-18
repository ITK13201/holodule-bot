package main

import (
	"encoding/json"
	"fmt"
	"github.com/ITK13201/holodule-bot/config"
	"github.com/ITK13201/holodule-bot/usecase/discord"
	"github.com/ITK13201/holodule-bot/usecase/video"
	"time"
)

func main() {
	videos := video.GetVideosOfToday()

	//=================
	// single text
	//=================
	now := time.Now().In(config.JST)
	content := discord.OnlyTextContent{
		Content: fmt.Sprintf("**[INFO] %s**: Holodule updated", now.Format(time.ANSIC)),
	}
	contentJson, _ := json.Marshal(content)

	discord.NotifyWithWebhook(contentJson, cfg.DiscordWebhookUrlDaily)

	//====================
	// notify videos
	//====================
	allEmbeds := []discord.Embed{}
	for i := 0; i < len(videos); i++ {
		allEmbeds = append(allEmbeds, discord.GetEmbed(videos[i]))
	}

	limit := len(allEmbeds)/10 + 1
	for i := 0; i < limit; i++ {
		x := i * 10
		var embeds []discord.Embed
		if i == limit-1 {
			embeds = allEmbeds[x:]
			if len(embeds) == 0 {
				continue
			}
		} else {
			embeds = allEmbeds[x : x+10]
		}

		content := discord.EmbedsContent{
			Embeds: embeds,
		}
		contentJson, _ := json.Marshal(content)

		discord.NotifyWIthBot(contentJson, cfg.DiscordChannelIdDaily)
	}
}
