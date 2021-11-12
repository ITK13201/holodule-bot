package main

import (
	"github.com/ITK13201/holodule-bot/usecase/discord"
	"github.com/ITK13201/holodule-bot/usecase/video"
)

func main() {
	videos := video.GetVideosOfToday()
	discord.Notify(videos)
}
