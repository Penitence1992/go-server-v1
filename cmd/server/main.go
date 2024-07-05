package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penitence1992/go-gin-server/internal/app"
	"github.com/penitence1992/go-gin-server/internal/errors"
	v1 "github.com/penitence1992/go-gin-server/internal/routers/v1"
	"github.com/penitence1992/go-gin-server/pkg/api"
	"github.com/penitence1992/go-gin-server/pkg/app/config"
	"github.com/penitence1992/go-gin-server/pkg/discovery"
	cerrors "github.com/penitence1992/go-gin-server/pkg/errors"
	"github.com/penitence1992/go-gin-server/pkg/server/starter"
	"github.com/penitence1992/go-gin-server/pkg/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	gitCommit  = "1"
	buildStamp = "1900-01-01"
)

func main() {

	log.Infoln("Server Starting --->")
	log.Infof("Git Commit : %s", gitCommit)
	log.Infof("Build Stamp : %s", buildStamp)

	setting, err := initConfig()
	if err != nil {
		panic(err)
	}

	// client, err := storage.Create(setting.CreateDatabaseSetting())
	utils.PanicIfNotNil(err)
	if setting.EurekaEnabled {
		err = discovery.DoRegistryAsync(10*time.Second, setting.Eureka)
		utils.PanicIfNotNil(err)
	}

	s := starter.NewServerStart(8080, "0.0.0.0", func(engine *gin.Engine) {
		utils.PanicIfNotNil(v1.RegistryHeadersRoute(engine))
	})
	s.AddErrorCapture(&errors.TestError{}, func(err error) *api.CwResponse {
		return api.Error(444, cerrors.ServerError, err.Error())
	})
	s.Start()
}

func initConfig() (*app.App, error) {
	c, err := config.GetCreator("", "app")
	if err != nil {
		return nil, err
	}
	var s = &app.App{}
	if err = c.GetConfig(s); err != nil {
		return nil, err
	}
	if err = s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}
