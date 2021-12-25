package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/ITK13201/holodule-bot/config"
)

type YoutubeSearchResult struct {
	VideoID      string
	VideoTitle   string
	ChannelID    string
	ChannelTitle string
}

func SearchYouTubeVideo(videoID string) []YoutubeSearchResult {
	cfg := *config.Cfg
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=snippet", videoID, cfg.YoutubeDataAPIKey)

	var resp *http.Response
	var err error

	resp, err = http.Get(url)
	if err != nil {
		log.Error("[!] " + err.Error())
	} else {
		log.Info("[*] " + resp.Status)
	}

	respTxt, _ := ioutil.ReadAll(resp.Body)
	// log.Info(string(respTxt))

	var respData map[string]interface{}
	err = json.Unmarshal(respTxt, &respData)
	if err != nil {
		log.Error(err.Error())
	}

	results := []YoutubeSearchResult{}
	respItems := respData["items"].([]interface{})
	for _, item := range respItems {
		assertedItem := item.(map[string]interface{})
		videoID := assertedItem["id"].(string)
		assertedItemSnippet := assertedItem["snippet"].(map[string]interface{})
		videoTitle := assertedItemSnippet["title"].(string)
		channelID := assertedItemSnippet["channelId"].(string)
		channelTitle := assertedItemSnippet["channelTitle"].(string)

		results = append(results, YoutubeSearchResult{
			VideoID:      videoID,
			VideoTitle:   videoTitle,
			ChannelID:    channelID,
			ChannelTitle: channelTitle,
		})
	}

	return results
}
