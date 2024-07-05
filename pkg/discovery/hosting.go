package discovery

import (
	"github.com/penitence1992/go-gin-server/pkg/discovery/configs"
	"github.com/penitence1992/go-gin-server/pkg/discovery/eureka"
	"github.com/penitence1992/go-gin-server/pkg/utils"
	"sync"
	"time"
)

func DoRegistryAsync(period time.Duration, eurekaConfig configs.EurekaConfig) error {
	nodes, err := eureka.CreateEurekaDiscovers(eurekaConfig)
	if err != nil {
		return err
	}
	taskGroup := &sync.WaitGroup{}
	for _, client := range nodes {
		rsignal := utils.CreateInterruptChan()
		executor := CreateTimerExecutor(client, rsignal, period)
		go func() {
			taskGroup.Add(1)
			defer taskGroup.Done()
			executor.Start()
		}()
	}
	return nil
}
