package processes

import (
	"github.com/ITK13201/holodule-bot/config"
	"github.com/ITK13201/holodule-bot/infrastructure"
	"github.com/ITK13201/holodule-bot/usecase/interactor"
)

var (
	distributorInteractor = interactor.NewDistributorInteractor(infrastructure.NewSqlHandler())
	videoInteractor       = interactor.NewVideoInteractor(infrastructure.NewSqlHandler())
	cfg                   = *config.Cfg
)
