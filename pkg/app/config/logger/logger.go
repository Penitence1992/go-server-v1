package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// init 这里的init方法主要是用于初始化设置logger日志的等级
func init() {
	lev := os.Getenv("loggerLevel")
	if l, err := log.ParseLevel(lev); err == nil {
		log.SetLevel(l)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
