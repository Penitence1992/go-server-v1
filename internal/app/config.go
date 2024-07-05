package app

import (
	"github.com/penitence1992/go-gin-server/pkg/discovery/configs"
	"github.com/penitence1992/go-gin-server/pkg/storage"
)

type App struct {
	DatabaseSetting storage.DatabaseSetting
	EurekaEnabled   bool
	Eureka          configs.EurekaConfig
}

func (d *App) Validate() (err error) {
	if err = d.DatabaseSetting.Validate(); err != nil {
		return
	}
	return nil
}
