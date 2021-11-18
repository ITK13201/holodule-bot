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

	discord.NotifyWithWebhook(contentJson)

	//====================
	// notify videos
	//====================
	allEmbedContents := []discord.EmbedContent{}
	for i := 0; i < len(videos); i++ {
		allEmbedContents = append(allEmbedContents, discord.GetEmbedContent(videos[i]))
	}

	limit := len(allEmbedContents)/10 + 1
	for i := 0; i < limit; i++ {
		x := i * 10
		var embedContents []discord.EmbedContent
		if i == limit-1 {
			embedContents = allEmbedContents[x:]
			if len(embedContents) == 0 {
				continue
			}
		} else {
			embedContents = allEmbedContents[x : x+10]
		}

		content := discord.Content{
			Embeds: embedContents,
		}
		contentJson, _ := json.Marshal(content)

		discord.NotifyWIthBot(contentJson)
	}
}
