package capture

import (
	"github.com/penitence1992/go-gin-server/pkg/api"
	"github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

type ErrorCapture func(err error) *api.CwResponse

type Catcher struct {
	handler map[reflect.Type]ErrorCapture
	lock    *sync.Mutex
}

func NewCatcher() *Catcher {
	return &Catcher{
		handler: make(map[reflect.Type]ErrorCapture),
		lock:    &sync.Mutex{},
	}
}

func (c *Catcher) RegisterErrorCapture(eType reflect.Type, f ErrorCapture) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.handler[eType] = f
}

func (c *Catcher) TryConvert(e error) (bool, *api.CwResponse) {
	t := reflect.TypeOf(e)
	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		logrus.Debugf("尝试转换类型为:%v的异常", t)
	}
	if h := c.handler[t]; h != nil {
		return true, h(e)
	}
	return false, nil
}
