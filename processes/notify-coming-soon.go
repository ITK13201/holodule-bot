package processes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ITK13201/holodule-bot/config"
	"github.com/ITK13201/holodule-bot/domain"
	"github.com/ITK13201/holodule-bot/usecase/discord"
	"github.com/ITK13201/holodule-bot/usecase/video"
	"log"
	"time"
)

func updateVideos() {
	videos := video.GetVideosWithDatetime()
	var distributorModel *domain.Distributor
	var err error

	for i := 0; i < len(videos); i++ {
		// update distributors
		distributorModel, err = distributorInteractor.GetByName(videos[i].Distributor.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				distributorModel, err = distributorInteractor.Add(domain.Distributor{
					Name:    videos[i].Distributor.Name,
					IconUrl: videos[i].Distributor.IconUrl,
				})
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		} else {
			err = distributorInteractor.Update(*distributorModel)
			if err != nil {
				log.Fatal(err)
			}
		}

		// update videos
		d := *distributorModel

		_, err = videoInteractor.GetBy3(d.Id, videos[i].Url, videos[i].Datetime)
		if err != nil {
			if err == sql.ErrNoRows {
				_, err = videoInteractor.Add(domain.Video{
					Distributor: d,
					Url:         videos[i].Url,
					Datetime:    videos[i].Datetime,
					ImageUrl:    videos[i].ImageUrl,
				})
				if err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		}
	}
}

func findVideoWithDatetime(vm domain.Video, videos []video.VideoWithDatetime) (*video.VideoWithDatetime, bool) {
	for i := 0; i < len(videos); i++ {
		v := videos[i]
		if vm.Distributor.Name == v.Distributor.Name && vm.Url == v.Url && vm.Datetime == v.Datetime {
			return &v, true
		}
	}
	return nil, false
}

func notify() {
	videos := video.GetVideosWithDatetime()
	videoModels, err := videoInteractor.GetComingSoon()
	if err != nil {
		if err == sql.ErrNoRows {
			return
		} else {
			log.Fatal(err)
		}
	}
	// utc to jst
	for i := 0; i < len(videoModels); i++ {
		videoModels[i].Utc2jst()
	}

	if len(videoModels) == 0 {
		return
	}

	//=================
	// single text
	//=================
	now := time.Now().In(config.JST)
	content := discord.OnlyTextContent{
		Content: fmt.Sprintf("**[INFO] %s**: Stream Coming Soon!", now.Format(time.ANSIC)),
	}
	contentJson, _ := json.Marshal(content)

	discord.NotifyWithWebhook(contentJson, cfg.DiscordWebhookUrlComingSoon)

	//====================
	// notify videos
	//====================
	for i := 0; i < len(videoModels); i++ {
		vm := videoModels[i]
		v, exists := findVideoWithDatetime(vm, videos)
		if !exists {
			continue
		}

		content := discord.EmbedContent{
			Embed: discord.GetEmbed(*v),
		}
		contentJson, _ := json.Marshal(content)

		discord.NotifyWIthBot(contentJson, cfg.DiscordChannelIdComingSoon)

		// add notified flags
		err := videoInteractor.UpdateNotifiedAt(vm.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NotifyComingSoon() {
	updateVideos()
	notify()
}
