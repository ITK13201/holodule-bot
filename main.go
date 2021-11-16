package main

import (
	"fmt"
	"github.com/ITK13201/holodule-bot/infrastructure"
	"github.com/ITK13201/holodule-bot/usecase/interactor"
)

func main() {
	//videos := video.GetVideosOfToday()
	//discord.Notify(videos)

	//d := domain.Distributor{
	//	Name: "laa",
	//	IconUrl: "lsls",
	//}
	distributorInteractor := interactor.NewDistributorInteractor(infrastructure.NewSqlHandler())
	distributor, _ := distributorInteractor.GetById(1)
	fmt.Println(*distributor)

	//v := domain.Video{
	//	Distributor: *distributor,
	//	Datetime: time.Now(),
	//}

	videoInteractor := interactor.NewVideoInteractor(infrastructure.NewSqlHandler())
	video, err := videoInteractor.GetById(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(video)
}
