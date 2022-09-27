package video

import (
	"github.com/ITK13201/holodule-bot/config"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const HoloduleUrl string = "https://schedule.hololive.tv/"

func getHoloduleHTML() *http.Response {
	url := HoloduleUrl
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

type Distributor struct {
	Name    string
	IconUrl string
}

type Video struct {
	Distributor Distributor
	Url         string
	Datetime    string
	ImageUrl    string
}

type VideoWithDatetime struct {
	Distributor Distributor
	Url         string
	Datetime    time.Time
	ImageUrl    string
}

func GetVideos() []Video {
	resp := getHoloduleHTML()
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	videos := []Video{}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	items := doc.Find("div.container > div.row > div.col-12.col-sm-12.col-md-12 > div.row > div.col-6.col-sm-4.col-md-3")
	items.Each(func(itemIdx int, itemSelection *goquery.Selection) {
		video := Video{}
		distributor := Distributor{}

		linkElem := itemSelection.Find("a.thumbnail")
		link, _ := linkElem.Attr("href")
		video.Url = link

		wrapper := itemSelection.Find("a.thumbnail > div.container > div.row")
		wrapper.Children().Each(func(wrapperIdx int, wrapperSelection *goquery.Selection) {
			if wrapperIdx == 0 {
				// top
				datetimeElem := wrapperSelection.Find("div.row.no-gutters > div.col-4.col-sm-4.col-md-4.text-left.datetime")
				datetime := datetimeElem.Text()
				video.Datetime = datetime
				nameElem := wrapperSelection.Find("div.row.no-gutters > div.col.text-right.name")
				distributor.Name = nameElem.Text()
			} else if wrapperIdx == 1 {
				// middle
				videoImgElem := wrapperSelection.Find("img")
				videoImgLink, _ := videoImgElem.Attr("src")
				video.ImageUrl = videoImgLink
			} else if wrapperIdx == 2 {
				// bottom
				iconElem := wrapperSelection.Find("div.row.no-gutters.justify-content-between > div.col.col-sm.col-md.col-lg.col-xl > img")
				iconLink, _ := iconElem.Attr("src")
				distributor.IconUrl = iconLink
			}
		})

		video.Distributor = distributor
		videos = append(videos, video)
	})

	for i := 0; i < len(videos); i++ {
		videos[i].Datetime = strings.TrimSpace(videos[i].Datetime)
		videos[i].Distributor.Name = strings.TrimSpace(videos[i].Distributor.Name)
	}

	return videos
}

func GetVideosWithDatetime() []VideoWithDatetime {
	videos := GetVideos()

	now := time.Now().In(config.JST)
	day := now.Day() - 1

	var prevHour = -1
	videosWithDatetime := []VideoWithDatetime{}
	for i := 0; i < len(videos); i++ {
		video := videos[i]
		datetimeStr := video.Datetime
		datetimeStrList := strings.Split(datetimeStr, ":")
		datetimeHour, _ := strconv.Atoi(datetimeStrList[0])
		datetimeMin, _ := strconv.Atoi(datetimeStrList[1])

		if prevHour > datetimeHour {
			day += 1
		}
		prevHour = datetimeHour

		datetime := time.Date(now.Year(), now.Month(), day, datetimeHour, datetimeMin, 0, 0, config.JST)

		videoWithDatetime := VideoWithDatetime{
			Distributor: Distributor{
				Name:    video.Distributor.Name,
				IconUrl: video.Distributor.IconUrl,
			},
			Url:      video.Url,
			Datetime: datetime,
			ImageUrl: video.ImageUrl,
		}
		videosWithDatetime = append(videosWithDatetime, videoWithDatetime)
	}

	return videosWithDatetime
}

func GetVideosOfToday() []VideoWithDatetime {
	videos := GetVideosWithDatetime()

	now := time.Now().In(config.JST)

	videosOfToday := []VideoWithDatetime{}
	for i := 0; i < len(videos); i++ {
		video := videos[i]
		if video.Datetime.Day() == now.Day() {
			videosOfToday = append(videosOfToday, video)
		}
	}

	return videosOfToday
}
