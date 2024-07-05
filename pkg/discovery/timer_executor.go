package discovery

import (
	"github.com/penitence1992/go-server-v1/pkg/discovery/api"
	"github.com/penitence1992/go-server-v1/pkg/discovery/runner"
	"github.com/penitence1992/go-server-v1/pkg/errors"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Executor struct {
	register   api.Discovery
	period     time.Duration
	sign       <-chan os.Signal
	interrupt  bool
	errorCount int
}

// CreateTimerExecutor 创建一个Eureka定时执行器
// register 为发现的实例
// sign 为接收中断信号
// period 更新的周期
func CreateTimerExecutor(register api.Discovery, sign <-chan os.Signal, period time.Duration) runner.Execute {
	return &Executor{
		register: register,
		period:   period,
		sign:     sign,
	}
}

func (e *Executor) Start() {
	ticker := time.NewTicker(e.period)
	defer ticker.Stop()
	defer e.doOnDone()
	go func() {
		select {
		case sg := <-e.sign:
			log.Infof("时间执行器接收到信号:%s", sg)
			e.interrupt = true
			return
		}
	}()

	for {
		e.errorCount = 0
		result := e.doRegisterAppInstance()
		// 不存在, 并且注册app失败, 则重新执行
		if !result {
			if e.interrupt {
				break
			}
			time.Sleep(5 * time.Second)
			continue
		} else {
			log.Info("注册实例成功")
		}
		log.Info("开始进行心跳发送")
		for {
			// 错误大于5次以后重新执行外层步骤
			if e.errorCount > 5 {
				break
			}
			if e.interrupt {
				return
			}
			select {
			case <-ticker.C:
				e.doOnHeartbeat()
			}

		}
	}
}

func (e *Executor) doRegisterAppInstance() bool {
	result, err := e.register.CreateInstance()
	if err != nil {
		log.Errorf("注册app实例失败: %v", err)
		return false
	}
	return result
}

func (e *Executor) doOnHeartbeat() {
	_, err := e.register.Heartbeat()
	if err != nil {
		switch e2 := err.(type) {
		case *errors.ProxyError:
			if e2.HttpStatus == 404 {
				log.Error("App未创建, 导致心跳发送失败, 重新创建App")
				e.errorCount += 100
				return
			}
		default:
			log.Errorf("心跳发送失败: %v", err)
			e.errorCount++
		}

	}
}

func (e *Executor) doOnDone() {
	_, err := e.register.RemoveInstance()
	if err != nil {
		log.Errorf("注册app失败: %v", err)
	}
}
